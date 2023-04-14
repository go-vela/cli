// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"fmt"

	"github.com/go-vela/cli/internal"
	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating worker configuration")

	// we need address for adding a worker
	if c.Action == internal.ActionAdd && len(c.Address) == 0 {
		return fmt.Errorf("no worker address provided")
	}

	// anything other than "get" action needs hostname
	if c.Action != internal.ActionGet && len(c.Hostname) == 0 {
		return fmt.Errorf("no worker hostname provided")
	}

	return nil
}
