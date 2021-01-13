// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package client

import (
	"fmt"
	"runtime"

	"github.com/go-vela/cli/internal"

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
	address := c.String(internal.FlagAPIAddress)

	// capture the token from the context
	token := c.String(internal.FlagAPIToken)

	// capture the access token from the context
	accessToken := c.String(internal.FlagAPIAccessToken)

	// capture the refresh token from the context
	refreshToken := c.String(internal.FlagAPIRefreshToken)

	// validate the provided configuration
	err := validate(address, token, accessToken, refreshToken)
	if err != nil {
		return nil, err
	}

	logrus.Tracef("creating Vela client for %s", address)

	// create the client id; will be in the form of
	// "vela; <version>; <os>; <architecture>"
	// used in user agent string in the sdk
	fmt.Printf("c.App: %+v", c.App)
	fmt.Printf("runtime.GOOS: %s", runtime.GOOS)
	fmt.Printf("runtime.GOARCH: %s", runtime.GOARCH)
	if len(c.App.Name) == 0 {
		fmt.Printf("c.App.Name: %s", c.App.Name)
	}
	fmt.Printf("c.App.Version: %s", c.App.Version)
	clientID := fmt.Sprintf("%s; %s; %s; %s",
		c.App.Name, c.App.Version, runtime.GOOS, runtime.GOARCH)

	// create a vela client from the provided address
	client, err := vela.NewClient(address, clientID, nil)
	if err != nil {
		return nil, err
	}

	logrus.Trace("setting token for Vela client")

	// pass the tokens to the client instance
	if len(accessToken) > 0 && len(refreshToken) > 0 {
		client.Authentication.SetAccessAndRefreshAuth(accessToken, refreshToken)
	}

	// pass the token to the client instance
	if len(token) > 0 {
		client.Authentication.SetTokenAuth(token)
	}

	return client, nil
}

// ParseEmptyToken digests the provided urfave/cli context
// and parses the provided configuration to produce a
// valid Vela client without token authentication.
func ParseEmptyToken(c *cli.Context) (*vela.Client, error) {
	logrus.Debug("parsing tokenless Vela client from provided configuration")

	// capture the address from the context
	address := c.String(internal.FlagAPIAddress)

	// check if client address is set
	if len(address) == 0 {
		return nil, fmt.Errorf("no client address provided")
	}

	logrus.Tracef("creating Vela client for %s", address)

	// create a vela client from the provided address
	return vela.NewClient(address, c.App.Name, nil)
}
