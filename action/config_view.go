// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/config"

	"github.com/urfave/cli/v2"
)

// ConfigView defines the command for inspecting the config file.
var ConfigView = &cli.Command{
	Name:        "config",
	Description: "Use this command to view the config file.",
	Usage:       "View the config file used in the CLI",
	Action:      configView,
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View the config file.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/config/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect the config file.
func configView(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action:   viewAction,
		File:     c.String("file"),
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
