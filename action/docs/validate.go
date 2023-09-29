// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating docs configuration")

	// check if multiple shells are set
	if c.Markdown && c.Man {
		return fmt.Errorf("multiple docs provided for generation")
	}

	// check if no shell is set
	if !c.Markdown && !c.Man {
		return fmt.Errorf("no docs provided for generation")
	}

	return nil
}
