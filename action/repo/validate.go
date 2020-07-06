// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating repo configuration")

	// check if repository action is not get
	if c.Action != "get" {
		// check if repository org is set
		if len(c.Org) == 0 {
			return fmt.Errorf("no repo org provided")
		}

		// check if repository name is set
		if len(c.Name) == 0 {
			return fmt.Errorf("no repo name provided")
		}
	}

	return nil
}
