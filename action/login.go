// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"
	"os"

	"github.com/go-vela/cli/action/login"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// Login defines the command for authenticating and logging in to Vela.
var Login = &cli.Command{
	Name:        "login",
	Description: "Use this command to authenticate and login to Vela.",
	Usage:       "Authenticate and login to Vela",
	Action:      runLogin,
	Flags: []cli.Flag{

		// API Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ADDR", "LOGIN_ADDR"},
			Name:    internal.FlagAPIAddress,
			Aliases: []string{"a"},
			Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
		},

		// User Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ACCESS_TOKEN", "CONFIG_ACCESS_TOKEN"},
			Name:    internal.FlagAPIAccessToken,
			Aliases: []string{"at"},
			Usage:   "access token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REFRESH_TOKEN", "CONFIG_REFRESH_TOKEN"},
			Name:    internal.FlagAPIRefreshToken,
			Aliases: []string{"rt"},
			Usage:   "refresh token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN"},
			Name:    internal.FlagAPIToken,
			Aliases: []string{"t"},
			Usage:   "token used for communication with the Vela server",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Login to Vela with terminal prompts.
    $ {{.HelpName}} --api.addr https://vela.example.com
  3. Login to Vela using a supplied Personal Access Token
    $ {{.HelpName}} --token foo

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/login/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
//  authenticate and login to Vela.
func runLogin(c *cli.Context) error {
	// load variables from the config file
	err := load(c)
	if err != nil {
		return err
	}

	// parse the Vela client from the context
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/client?tab=doc#ParseEmptyToken
	client, err := client.ParseEmptyToken(c)
	if err != nil {
		return err
	}

	// create the login configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config
	l := &login.Config{}

	err = l.PromptBrowserConfirm(os.Stdin, c.String(internal.FlagAPIAddress))
	if err != nil {
		return err
	}

	// execute the login call for the login configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.Login
	err = l.Login(client, c.String(internal.FlagAPIAddress))
	if err != nil {
		return err
	}

	// no error means above means we have tokens, set them
	err = c.Set(internal.FlagAPIAccessToken, l.AccessToken)
	if err != nil {
		return err
	}

	err = c.Set(internal.FlagAPIRefreshToken, l.RefreshToken)
	if err != nil {
		return err
	}

	// ask to write config
	err = l.PromptConfigConfirm(os.Stdin)
	if err != nil {
		return err
	}

	// generate new config
	// ideally, we update the config
	// but if one doesn't exist, it
	// currently errors out
	err = configGenerate(c)
	if err != nil {
		return err
	}

	logrus.Info("configuration successfully created - enjoy")
	// err = configUpdate(c)
	// if err != nil {
	// 	return err
	// }

	return nil
}
