// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/dashboard"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
)

// CommandAdd defines the command for creating a dashboard.
var CommandAdd = &cli.Command{
	Name:        "dashboard",
	Description: "Use this command to add a dashboard.",
	Usage:       "Add a new dashboard from the provided configuration",
	Action:      add,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_DASHBOARD_NAME", "DASHBOARD_NAME"},
			Name:    "name",
			Usage:   "provide the name for the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_REPOS", "DASHBOARD_REPOS"},
			Name:    "repos",
			Usage:   "provide the list of repositories for the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_REPOS_BRANCH", "DASHBOARD_REPOS_BRANCH"},
			Name:    "branches",
			Aliases: []string{"branch"},
			Usage:   "filter builds in all repositories by branch",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_REPOS_EVENT", "DASHBOARD_REPOS_EVENT"},
			Name:    "events",
			Aliases: []string{"event"},
			Usage:   "filter builds in all repositories by event",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_ADMINS", "DASHBOARD_ADMINS"},
			Name:    "admins",
			Usage:   "provide the list of admins for the dashboard",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "REPO_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Add a repository with push and pull request enabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --event push --event pull_request
  2. Add a repository with all event types enabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --event push --event pull_request --event tag --event deployment --event comment
  3. Add a repository with a longer build timeout.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --timeout 90
  4. Add a repository when config or environment variables are set.
    $ {{.HelpName}} --event push --event pull_request
  5. Add a repository with a starting build number.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --counter 90
  6. Add a repository with a starlark pipeline file.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --pipeline-type starlark
  7. Add a repository with approve build setting set to fork-no-write.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --approve-build fork-no-write

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/repo/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to create a repo.
func add(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// parse the Vela client from the context
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/client?tab=doc#Parse
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// create the dashboard configuration
	d := &dashboard.Config{
		Action:    internal.ActionAdd,
		Name:      c.String("name"),
		AddRepos:  c.StringSlice("repos"),
		Branches:  c.StringSlice("branches"),
		Events:    c.StringSlice("events"),
		AddAdmins: c.StringSlice("admins"),
		Output:    c.String(internal.FlagOutput),
	}

	// validate dashboard configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the dashboard configuration
	return d.Add(client)
}
