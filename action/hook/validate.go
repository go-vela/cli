// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

import (
	"fmt"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	// check if hook org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no hook org provided")
	}

	// check if hook repo is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no hook repo provided")
	}

	// check if hook action is view
	if c.Action == "view" {
		// check if hook number is set
		if c.Number <= 0 {
			return fmt.Errorf("no hook number provided")
		}
	}

	return nil
}
