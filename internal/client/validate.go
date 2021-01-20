// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package client

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// validate is a helper function to verify the input provided.
func validate(address, token, accessToken, refreshToken string) error {
	logrus.Debug("validating provided configuration for Vela client")

	// check if client address is set
	if len(address) == 0 {
		return fmt.Errorf("no client address provided")
	}

	// check that a token is set
	if len(token) == 0 && (len(accessToken) == 0 && len(refreshToken) == 0) {
		return fmt.Errorf("no client token provided")
	}

	return nil
}
