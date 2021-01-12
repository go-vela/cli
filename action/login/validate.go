// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating login configuration")

	// check if address is set
	if len(c.Address) == 0 {
		return fmt.Errorf("no address provided")
	}

	// check if address is right format
	_, err := url.ParseRequestURI(c.Address)
	if err != nil {
		return err
	}

	return nil
}
