// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating login configuration")

	return nil
}
