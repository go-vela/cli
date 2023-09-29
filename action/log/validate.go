// SPDX-License-Identifier: Apache-2.0

package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating log configuration")

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

	return nil
}
