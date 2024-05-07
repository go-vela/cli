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

// CommandGet defines the command for viewing all user dashboards.
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
			Usage:   "output the repo and build information for the dashboard",
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
  1. Get user dashboards.
    $ {{.HelpName}}
  2. Get user dashboards with repo and build information.
    $ {{.HelpName}} --full

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/dashboard/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to get user dashboards.
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

	// execute the get call for the dashboard configuration
	return d.Get(client)
}
