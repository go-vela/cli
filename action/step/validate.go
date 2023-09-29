// SPDX-License-Identifier: Apache-2.0

package step

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating step configuration")

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
