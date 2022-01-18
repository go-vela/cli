// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"

	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal"

	"github.com/urfave/cli/v2"
)

// CommandView defines the command for inspecting the config file.
var CommandView = &cli.Command{
	Name:        "config",
	Description: "Use this command to view the config file.",
	Usage:       "View the config file used in the CLI",
	Action:      view,
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
func view(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action: internal.ActionView,
		File:   c.String(internal.FlagConfig),
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
