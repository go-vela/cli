// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"fmt"
	"net/http"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/library"

	"github.com/sirupsen/logrus"
)

// Config represents the configuration necessary
// to perform login related requests with Vela.
type Config struct {
	Action   string
	Username string
	Password string
	OTP      string
	Retry    bool
}

// Login authenticates and logs in to Vela via the API based off the provided configuration.
func (c *Config) Login(client *vela.Client) error {
	logrus.Debug("executing login for login configuration")

	// create the login object
	//
	// https://pkg.go.dev/github.com/go-vela/types/library?tab=doc#Login
	l := &library.Login{
		Username: vela.String(c.Username),
		Password: vela.String(c.Password),
		OTP:      vela.String(c.OTP),
	}

	logrus.Tracef("starting login for user %s", c.Username)

	// send API call to login to Vela
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#AuthorizationService.Login
	auth, resp, err := client.Authorization.Login(l)
	if err != nil {
		// log the response object for troubleshooting purposes
		logrus.Debug(resp)

		// check if expected unauthorized (401) status code was returned
		if resp.StatusCode != http.StatusUnauthorized {
			// log a warning indicating the unexpected status code received
			logrus.Warningf("unexpected status code received: %d", resp.StatusCode)

			// return the error received
			return err
		}
	}

	// check if an unexpected status code was returned
	if resp.StatusCode > http.StatusUnauthorized {
		// log the response object for troubleshooting purposes
		logrus.Debug(resp)

		// return the error received
		return err
	}

	// check if the token is empty from the response
	if len(auth.GetToken()) == 0 {
		logrus.Trace("no token value found in response")

		// check if the retry mechanism is false indicating
		// we need to send another request with a OTP
		if !c.Retry {
			// set the retry mechanism to true
			c.Retry = true

			// return an error indicating we need to retry with OTP
			return fmt.Errorf("no token received - retrying with one time password (OTP)")
		}

		// return an error indicating the inability to retrieve a token
		return fmt.Errorf("unable to retrieve token from login request")
	}

	// output the message in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(fmt.Sprintf("token: %s", auth.GetToken()))
}
