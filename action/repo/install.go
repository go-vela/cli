// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"fmt"
	"strconv"

	"github.com/cli/browser"
	"github.com/sirupsen/logrus"

	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// Install executes the repo app installation process, which should redirect to the SCM web flow.
func (c *Config) Install(client *vela.Client, repo *api.Repo) error {
	logrus.Debug("executing app install for repo configuration")

	// start the local server
	err := c.StartServer()
	if err != nil {
		return err
	}

	// request the install URL from the server
	installHTMLURL, _, err := client.Repo.InstallHTMLURL(repo.GetOrg(), repo.GetName())
	if err != nil {
		return err
	}

	// attach contextual information like cli type and local server port
	*installHTMLURL = fmt.Sprintf(
		"%s&type=%s&port=%s",
		*installHTMLURL,
		"cli", strconv.Itoa(c.server.Port()),
	)

	// launch the login process in the browser
	err = browser.OpenURL(*installHTMLURL)
	if err != nil {
		return err
	}

	// capture result from local server
	err = c.WaitForResult(client)
	if err != nil {
		return err
	}

	return nil
}

// WaitForResult will wait for the callback and handle the response.
func (c *Config) WaitForResult(client *vela.Client) error {
	logrus.Debug("waiting for app installation server callback")

	// waiting for local server to receive the redirect
	_, err := c.server.WaitForResult()
	if err != nil {
		return err
	}

	return nil
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
