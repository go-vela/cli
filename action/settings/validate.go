// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating settings configuration")

	return nil
}
