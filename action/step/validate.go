// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"fmt"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	// check if step org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no step org provided")
	}

	// check if step repo is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no step repo provided")
	}

	// check if step build is set
	if c.Build <= 0 {
		return fmt.Errorf("no step build provided")
	}

	// check if step action is view
	if c.Action == "view" {
		// check if step number is set
		if c.Number <= 0 {
			return fmt.Errorf("no step number provided")
		}
	}

	return nil
}
