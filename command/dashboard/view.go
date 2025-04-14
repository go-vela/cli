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

// CommandView defines the command for viewing a dashboard.
var CommandView = &cli.Command{
	Name:        "dashboard",
	Description: "Use this command to view a dashboard.",
	Usage:       "View a dashboard from the provided configuration",
	Action:      view,
	Flags: []cli.Flag{

		// Dashboard Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_DASHBOARD_ID", "DASHBOARD_ID"),
			Name:    "id",
			Usage:   "provide the uuid for the dashboard",
		},

		// Output Flags

		&cli.BoolFlag{
			Sources: cli.EnvVars("VELA_FULL", "DASHBOARD_FULL"),
			Name:    "full",
			Usage:   "output the repo and build information for the dashboard",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "REPO_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View a dashboard.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8
  2. View a dashboard with repo and build information.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8 --full

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/dashboard/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to view a dashboard.
func view(ctx context.Context, c *cli.Command) error {
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
		Action: internal.ActionView,
		ID:     c.String("id"),
		Full:   c.Bool("full"),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate dashboard configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the dashboard configuration
	return d.View(client)
}
