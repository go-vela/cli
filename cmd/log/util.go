// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"github.com/go-vela/cli/util"
	"github.com/urfave/cli/v2"
)

// helper function to load global configuration if set
// via config or environment and validate the user input in the command
func validate(c *cli.Context) error {

	// load configuration
	if len(c.String("org")) == 0 {
		c.Set("org", c.String("org"))
	}
	if len(c.String("repo")) == 0 {
		c.Set("repo", c.String("repo"))
	}

	// validate the user input in the command
	if len(c.String("org")) == 0 {
		return util.InvalidCommand("org")
	}
	if len(c.String("repo")) == 0 {
		return util.InvalidCommand("repo")
	}
	if c.Int("build-number") == 0 {
		return util.InvalidCommand("build-number")
	}

	if len(c.String("type")) == 0 {
		return util.InvalidCommand("type")
	} else if c.String("type") != "step" && c.String("type") != "service" {
		return util.InvalidFlag("type")
	}

	return nil
}
