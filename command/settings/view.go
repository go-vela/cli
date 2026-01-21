// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/settings"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandView defines the command for inspecting the platform settings record.
var CommandView = &cli.Command{
	Name:        "settings",
	Description: "Use this command to view platform settings.",
	Usage:       "View details for platform settings",
	Aliases:     []string{"platform"},
	Action:      view,
	Flags: []cli.Flag{

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "SETTINGS_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View platform settings.
    $ {{.FullName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/settings/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect.
func view(ctx context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// parse the Vela client from the context
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// create the configuration
	s := &settings.Config{
		Action: internal.ActionView,
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate settings configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/settings?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the settings configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/settings?tab=doc#Config.View
	return s.View(ctx, client)
}
