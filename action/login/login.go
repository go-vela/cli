// SPDX-License-Identifier: Apache-2.0

package login

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cli/browser"
	"github.com/sirupsen/logrus"

	"github.com/go-vela/sdk-go/vela"
)

// Config represents the configuration necessary
// to perform login related requests with Vela.
type Config struct {
	server *localServer

	Action       string
	Address      string
	AccessToken  string
	RefreshToken string
}

// Login authenticates and logs in to Vela via the API based off the provided configuration.
func (c *Config) Login(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing login for login configuration")

	// start the local server
	err := c.StartServer(ctx)
	if err != nil {
		return err
	}

	// create the options object for the client
	opts := &vela.LoginOptions{
		Type: "cli",
		Port: strconv.Itoa(c.server.Port()),
	}

	// get the login url to use for the
	// browser session
	url, err := client.Authorization.GetLoginURL(opts)
	if err != nil {
		return err
	}

	logrus.Tracef("got login url: %s", url)

	// launch the login process in the browser
	err = browser.OpenURL(url)
	if err != nil {
		return err
	}

	// capture the tokens
	err = c.Tokens(ctx, client, c.Address)
	if err != nil {
		return err
	}

	return nil
}

// Tokens will wait for the callback and make a request
// to exchange the the callback payload for an
// access token.
func (c *Config) Tokens(ctx context.Context, client *vela.Client, addr string) error {
	logrus.Debug("waiting for tokens")

	// waiting for local server to receive the redirect
	code, err := c.server.WaitForCode()
	if err != nil {
		return err
	}

	logrus.Debug("tokens received")

	// prepare options to exchange for token
	opt := &vela.OAuthExchangeOptions{
		Code:  code.Code,
		State: code.State,
	}

	// send API call to exchange tokens to Vela
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#AuthenticationService.ExchangeTokens
	at, rt, resp, err := client.Authentication.ExchangeTokens(ctx, opt)
	if err != nil {
		// log the response object for troubleshooting purposes
		logrus.Debug(resp)

		if resp == nil {
			return fmt.Errorf("unable to successfully connect to %s: %w", addr, err)
		}

		// check if expected unauthorized (401) status code was returned
		if resp.StatusCode != http.StatusUnauthorized {
			// log a warning indicating the unexpected status code receive
			logrus.Warningf("unexpected status code received: %d", resp.StatusCode)

			// return the error received
			return err
		}
	}

	if len(at) == 0 || len(rt) == 0 {
		logrus.Trace("no token value found in response")

		return fmt.Errorf("unable to retrieve tokens from authentication request")
	}

	// set the access token
	c.AccessToken = at

	// set the refresh cookie
	c.RefreshToken = rt

	return nil
}

// StartServer starts a local server as part of the
// auth flow. It will handle the callback.
func (c *Config) StartServer(ctx context.Context) error {
	logrus.Debug("starting local server")

	// set up the local server to capture the redirect from auth
	server, err := bindLocalServer(ctx)
	if err != nil {
		return err
	}

	logrus.Debug("local server is bound")

	// store on struct
	c.server = server

	// start the server up
	go func() {
		_ = c.server.Serve()
	}()

	logrus.Debug("local server started")

	return nil
}
