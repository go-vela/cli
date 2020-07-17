// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating login configuration")

	// check if login username is set
	if len(c.Username) == 0 {
		return fmt.Errorf("no login username provided")
	}

	// check if login password is set
	if len(c.Password) == 0 {
		return fmt.Errorf("no login password provided")
	}

	// check if the retry mechanism is configured
	if c.Retry {
		// check if login OTP is set
		if len(c.OTP) == 0 {
			return fmt.Errorf("no login one time password (OTP) provided")
		}
	}

	return nil
}
