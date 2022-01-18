// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package client

import (
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
)

// validate is a helper function to verify the input provided.
func validate(address, token, accessToken, refreshToken string) error {
	logrus.Debug("validating provided configuration for Vela client")

	// check if client address is set
	if len(address) == 0 {
		return fmt.Errorf("no client address provided")
	}

	// check for valid URL
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return fmt.Errorf("client address is not a valid url")
	}

	// check that a token is set
	if len(token) == 0 && (len(accessToken) == 0 && len(refreshToken) == 0) {
		return fmt.Errorf("no client token provided")
	}

	return nil
}
