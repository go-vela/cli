// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating deployment configuration")

	// check if deployment org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no deployment org provided")
	}

	// check if deployment repo is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no deployment repo provided")
	}

	// check if deployment action is add
	if c.Action == "add" {
		// check if deployment ref is set
		if len(c.Ref) == 0 {
			logrus.Debug("no ref provided. Using repo default branch...")
		}

		// check if deployment target is set
		if len(c.Target) == 0 {
			logrus.Debug("no deployment target provided")
		}
	}

	// check if deployment action is view
	if c.Action == "view" {
		// check if deployment number is set
		if c.Number <= 0 {
			return fmt.Errorf("no deployment number provided")
		}
	}

	return nil
}
