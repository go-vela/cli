// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"fmt"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	// check if log org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no log org provided")
	}

	// check if log repo is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no log repo provided")
	}

	// check if log build is set
	if c.Build <= 0 {
		return fmt.Errorf("no log build provided")
	}

	// check if log action is view
	if c.Action == "view" {
		// check if service or step is set
		if c.Service <= 0 && c.Step <= 0 {
			return fmt.Errorf("no log service or step number provided")
		}
	}

	return nil
}
