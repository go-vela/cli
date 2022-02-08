// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"fmt"

	"github.com/go-vela/cli/internal"
	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating build configuration")

	// check if build org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no build org provided")
	}

	// check if build repo is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no build repo provided")
	}

	// check if build action is restart or view
	if c.Action == internal.ActionRestart || c.Action == internal.ActionView {
		// check if build number is set
		if c.Number <= 0 {
			return fmt.Errorf("no build number provided")
		}
	}

	return nil
}
