// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"github.com/go-vela/cli/util"
	"github.com/urfave/cli/v2"
)

var logTypes = [...]string{"step", "service"}

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

	logType := c.String("type")
	if len(logType) == 0 {
		return util.InvalidCommand("type")
	} else if !IsValidType(logType) {
		return util.InvalidFlagValue(logType, "type")
	}

	return nil
}

func IsValidType(givenType string) bool {
	for _, validType := range logTypes {
		if validType == givenType {
			return true
		}
	}
	return false
}
