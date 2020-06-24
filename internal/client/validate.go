// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package client

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// validate is a helper function to verify the input provided.
func validate(address, token string) error {
	logrus.Debug("validating provided configuration for Vela client")

	// check if client address is set
	if len(address) == 0 {
		return fmt.Errorf("no client address provided")
	}

	// check if client token is set
	if len(token) == 0 {
		return fmt.Errorf("no client token provided")
	}

	return nil
}
