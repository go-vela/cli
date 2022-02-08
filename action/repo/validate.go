// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"

	"github.com/go-vela/cli/internal"
	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating repo configuration")

	// check if repository action is not get
	if c.Action != internal.ActionGet {
		// check if repository org is set
		if len(c.Org) == 0 {
			return fmt.Errorf("no repo org provided")
		}

		// check if repository action is not syncAll
		if c.Action != internal.ActionSyncAll {
			// check if repository name is set
			if len(c.Name) == 0 {
				return fmt.Errorf("no repo name provided")
			}
		}
	}

	return nil
}
