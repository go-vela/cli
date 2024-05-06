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

// CommandGet defines the command for viewing a dashboard.
var CommandGet = &cli.Command{
	Name:        "dashboard",
	Aliases:     []string{"dashboards"},
	Description: "Use this command to get user dashboards.",
	Usage:       "Get user dashboards from the provided configuration",
	Action:      get,
	Flags: []cli.Flag{
		// Output Flags
		&cli.BoolFlag{
			EnvVars: []string{"VELA_FULL", "DASHBOARD_FULL"},
			Name:    "full",
			Usage:   "output the full details of the dashboard",
		},

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "DASHBOARD_OUTPUT"},
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
// and create the object used to view a dashboard.
func get(c *cli.Context) error {
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
		Action: internal.ActionGet,
		Full:   c.Bool("full"),
		Output: c.String(internal.FlagOutput),
	}

	// validate dashboard configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the dashboard configuration
	return d.Get(client)
}
