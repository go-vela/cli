// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"

	"github.com/cli/browser"
)

// Config represents the configuration necessary
// to perform login related requests with Vela.
type Config struct {
	server *localServer

	Action              string
	Address             string
	PersonalAccessToken string
	AccessToken         string
	RefreshToken        string
}

// StartServer starts a local server as part of the
// auth flow. It will handle the callback.
func (c *Config) StartServer() error {
	logrus.Debug("starting local server")

	// set up the local server to capture the redirect from auth
	server, err := bindLocalServer()
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

// Tokens will wait for the callback and make a request
// to exchange the the callback payload for an
// access token.
func (c *Config) Tokens(addr string) error {
	logrus.Debug("waiting for tokens")

	// waiting for local server to receive the redirect
	code, err := c.server.WaitForCode()
	if err != nil {
		return err
	}

	// prep to send the next request to exchange for tokens
	query := url.Values{}
	query.Set("state", code.State)
	query.Set("code", code.Code)

	// TODO: this should be in the SDK
	url := fmt.Sprintf("%s?%s", fmt.Sprintf("%s/authenticate", addr), query.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// create a new http client with timeout
	httpClient := http.DefaultClient
	httpClient.Timeout = 15 * time.Second

	// send the request to get tokens
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// the refresh token will be in a cookie in the response
	rt := extractRefreshToken(resp.Cookies())

	// set the refresh cookie (not checking if it's empty)
	c.RefreshToken = rt

	// capture the token returned in the JSON response
	l := &library.Login{}
	err = json.NewDecoder(resp.Body).Decode(l)
	if err != nil {
		return err
	}

	// set the access token
	c.AccessToken = l.GetToken()

	return nil
}

// Login authenticates and logs in to Vela via the API based off the provided configuration.
func (c *Config) Login(client *vela.Client) error {
	logrus.Debug("executing login for login configuration")

	// start the local server
	err := c.StartServer()
	if err != nil {
		return err
	}

	// create the options object for the client
	opts := &vela.LoginOpts{
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
	err = c.Tokens(c.Address)
	if err != nil {
		return err
	}

	return nil
}

// extractRefreshToken is a helper function to extract
// the refresh token from the supplied cookie slice.
func extractRefreshToken(cookies []*http.Cookie) string {
	c := ""

	// loop over the cookies to find the refresh cookie
	for _, cookie := range cookies {
		if cookie.Name == constants.RefreshTokenName {
			c = cookie.Value
		}
	}

	return c
}
