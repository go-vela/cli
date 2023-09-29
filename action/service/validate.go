// SPDX-License-Identifier: Apache-2.0

package service

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating service configuration")

	// check if service org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no service org provided")
	}

	// check if service repo is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no service repo provided")
	}

	// check if service build is set
	if c.Build <= 0 {
		return fmt.Errorf("no service build provided")
	}

	// check if service action is view
	if c.Action == "view" {
		// check if service number is set
		if c.Number <= 0 {
			return fmt.Errorf("no service number provided")
		}
	}

	return nil
}
