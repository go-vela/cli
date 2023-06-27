// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package schedule

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating schedule configuration")

	// check if schedule org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no schedule org provided")
	}

	// check if schedule repo is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no schedule repo provided")
	}

	// check if schedule action is not get
	if c.Action != "get" {
		// check if schedule name is set
		if len(c.Name) == 0 {
			return fmt.Errorf("no schedule name provided")
		}
	}

	// check if schedule action is add
	if c.Action == "add" {
		// check if schedule entry is set
		if len(c.Entry) == 0 {
			return fmt.Errorf("no schedule entry provided")
		}
	}

	return nil
}
