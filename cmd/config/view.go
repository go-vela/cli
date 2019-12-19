// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

// ViewCmd defines the command for viewing a configuration file.
var ViewCmd = cli.Command{
	Name:        "config",
	Description: "Use this command to view a config file.",
	Usage:       "View details of the provided config file",
	Action:      view,
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. View CLI config.
    $ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to execute a generate repo cli command
func view(c *cli.Context) error {
	data, err := ioutil.ReadFile(c.GlobalString("config"))
	if err != nil {
		return fmt.Errorf("unable to read yaml config file: %v", err)
	}

	fmt.Println(string(data))

	return nil
}
