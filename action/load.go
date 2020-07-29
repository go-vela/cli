// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"github.com/go-vela/cli/action/config"

	"github.com/urfave/cli/v2"
)

// load is a helper function that loads the necessary configuration for the CLI.
func load(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action: loadAction,
		File:   c.String("config"),
	}

	// validate config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Validate
	err := conf.Validate()
	if err == nil {
		// execute the load call for the config file configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Load
		err = conf.Load(c)
		if err != nil {
			return err
		}
	}

	return nil
}
