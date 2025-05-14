// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/dashboard"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandAdd defines the command for creating a dashboard.
var CommandAdd = &cli.Command{
	Name:        "dashboard",
	Description: "Use this command to add a dashboard.",
	Usage:       "Add a new dashboard from the provided configuration",
	Action:      add,
	Flags: []cli.Flag{

		// Dashboard Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_DASHBOARD_NAME", "DASHBOARD_NAME"),
			Name:    "name",
			Usage:   "provide the name for the dashboard",
		},
		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_DASHBOARD_REPOS", "DASHBOARD_REPOS"),
			Name:    "repos",
			Aliases: []string{"add-repos"},
			Usage:   "provide the list of repositories (org/repo) for the dashboard",
		},
		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_DASHBOARD_REPOS_BRANCH", "DASHBOARD_REPOS_BRANCH"),
			Name:    "branches",
			Aliases: []string{"branch"},
			Usage:   "filter builds in all repositories by branch",
		},
		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_DASHBOARD_REPOS_EVENT", "DASHBOARD_REPOS_EVENT"),
			Name:    "events",
			Aliases: []string{"event"},
			Usage:   "filter builds in all repositories by event",
		},
		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_DASHBOARD_ADMINS", "DASHBOARD_ADMINS"),
			Name:    "admins",
			Usage:   "provide the list of admins for the dashboard",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "REPO_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Add a dashboard.
    $ {{.HelpName}} --name my-dashboard
  2. Add a dashboard with repositories.
    $ {{.HelpName}} --name my-dashboard --repos Org-1/Repo-1,Org-2/Repo-2
  3. Add a dashboard with repositories filtering builds by pushes to main.
    $ {{.HelpName}} --name my-dashboard --repos Org-1/Repo-1,Org-2/Repo-2 --branch main --event push
  4. Add a dashboard with multiple admins.
    $ {{.HelpName}} --name my-dashboard --admins JohnDoe,JaneDoe

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/dashboard/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to create a dashboard.
func add(_ context.Context, c *cli.Command) error {
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
		Color:     output.ColorOptionsFromCLIContext(c),
	}

	// validate dashboard configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the dashboard configuration
	return d.Add(client)
}
