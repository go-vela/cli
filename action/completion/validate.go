// SPDX-License-Identifier: Apache-2.0

package completion

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating completion configuration")

	// check if multiple shells are set
	if c.Bash && c.Zsh {
		return fmt.Errorf("multiple shells provided for completion")
	}

	// check if no shell is set
	if !c.Bash && !c.Zsh {
		return fmt.Errorf("no shell provided for completion")
	}

	return nil
}
