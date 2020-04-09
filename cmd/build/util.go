// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/urfave/cli/v2"
)

// helper function to load global configuration if set
// via config or environment and validate the user input in the command
func validate(c *cli.Context) error {
	// load configuration
	if len(c.String("org")) == 0 {
		err := c.Set("org", c.String("org"))
		if err != nil {
			return fmt.Errorf("unable to set context: %w", err)
		}
	}

	if len(c.String("repo")) == 0 {
		err := c.Set("repo", c.String("repo"))
		if err != nil {
			return fmt.Errorf("unable to set context: %w", err)
		}
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
