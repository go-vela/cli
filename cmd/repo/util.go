// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"github.com/go-vela/cli/util"
	"github.com/urfave/cli"
)

// helper function to load global configuration if set
// via config or environment and validate the user input in the command
func validate(c *cli.Context) error {

	// load configuration
	if len(c.String("org")) == 0 {
		c.Set("org", c.GlobalString("org"))
	}
	if len(c.String("repo")) == 0 {
		c.Set("repo", c.GlobalString("repo"))
	}

	// validate the user input in the command
	if len(c.String("org")) == 0 {
		return util.InvalidCommand("org")
	}
	if len(c.String("repo")) == 0 {
		return util.InvalidCommand("repo")
	}

	return nil
}
