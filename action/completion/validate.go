// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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
