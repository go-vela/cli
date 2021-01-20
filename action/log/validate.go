// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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
