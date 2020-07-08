// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating config file configuration")

	// check if config file is set
	if len(c.File) == 0 {
		return fmt.Errorf("no config file provided")
	}

	// check if config file exists
	_, err := os.Stat(c.File)
	if err != nil {
		// check if a not exist err was returned
		if os.IsNotExist(err) {
			return fmt.Errorf("no config file found @ %s", c.File)
		}

		return err
	}

	return nil
}
