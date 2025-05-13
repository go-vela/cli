// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal"
)

// CommandView defines the command for inspecting the config file.
var CommandView = &cli.Command{
	Name:        "config",
	Description: "Use this command to view the config file.",
	Usage:       "View the config file used in the CLI",
	Action:      view,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Hidden: true,
			Name:   "fs.mem-map",
			Usage:  "use memory mapped files for the config file (for testing)",
			Value:  false,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View the config file.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/config/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect the
// config file.
func view(_ context.Context, c *cli.Command) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action:    internal.ActionView,
		File:      c.String(internal.FlagConfig),
		UseMemMap: c.Bool("fs.mem-map"),
	}

	// validate config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Validate
	err := conf.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.View
	return conf.View()
}
