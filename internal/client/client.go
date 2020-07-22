// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package client

import (
	"fmt"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// Parse digests the provided urfave/cli context
// and parses the provided configuration to
// produce a valid Vela client.
func Parse(c *cli.Context) (*vela.Client, error) {
	logrus.Debug("parsing Vela client from provided configuration")

	// capture the address from the context
	address := c.String(KeyAddress)

	// capture the token from the context
	token := c.String(KeyToken)

	// validate the provided configuration
	err := validate(address, token)
	if err != nil {
		return nil, err
	}

	logrus.Tracef("creating Vela client for %s", address)

	// create a vela client from the provided address
	client, err := vela.NewClient(address, nil)
	if err != nil {
		return nil, err
	}

	logrus.Trace("setting token for Vela client")

	// set the authentication mechanism from the provided token
	client.Authentication.SetTokenAuth(token)

	return client, nil
}

// ParseEmptyToken digests the provided urfave/cli context
// and parses the provided configuration to produce a
// valid Vela client without token authentication.
func ParseEmptyToken(c *cli.Context) (*vela.Client, error) {
	logrus.Debug("parsing tokenless Vela client from provided configuration")

	// capture the address from the context
	address := c.String(KeyAddress)

	// check if client address is set
	if len(address) == 0 {
		return nil, fmt.Errorf("no client address provided")
	}

	logrus.Tracef("creating Vela client for %s", address)

	// create a vela client from the provided address
	return vela.NewClient(address, nil)
}
