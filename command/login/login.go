// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"fmt"
	"os"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/action/login"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// CommandLogin defines the command for authenticating and logging in to Vela.
var CommandLogin = &cli.Command{
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
		&cli.BoolFlag{
			EnvVars: []string{"VELA_YES_ALL", "CONFIG_YES_ALL"},
			Name:    "yes-all",
			Aliases: []string{"y"},
			Usage:   "auto-confirm all prompts",
			Value:   false,
		},

		// the following flag is only present to help clear
		// existing legacy tokens
		&cli.StringFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN"},
			Name:    internal.FlagAPIToken,
			Usage:   "hidden flag to clear existing token",
			Hidden:  true,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Login to Vela (will launch browser).
    $ {{.HelpName}} --api.addr https://vela.example.com

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/login/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to authenticate
// and login to Vela.
func runLogin(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
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
	l := &login.Config{
		Address: c.String(internal.FlagAPIAddress),
	}

	// show a prompt to open a browser, unless yes-all flag is set
	if !c.Bool("yes-all") {
		// prompt user to confirm opening browser
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.PromptBrowserConfirm
		err = l.PromptBrowserConfirm(os.Stdin)
		if err != nil {
			return err
		}
	}

	// execute the login call for the login configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.Login
	err = l.Login(client)
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

	// show a prompt to write config, unless yes-all flag is set
	if !c.Bool("yes-all") {
		// prompt user to confirm writing new config
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.PromptConfigConfirm
		err = l.PromptConfigConfirm(os.Stdin)
		if err != nil {
			logrus.Warn("configuration not saved")
			return err
		}
	}

	// remove existing token from the config
	// before writing
	err = c.Set(internal.FlagAPIToken, "")
	if err != nil {
		// fail silently; not returning err
		logrus.Debugf("error while emptying token: %v", err)
	}

	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action:       internal.ActionGenerate,
		File:         c.String(internal.FlagConfig),
		Addr:         c.String(internal.FlagAPIAddress),
		Token:        c.String(internal.FlagAPIToken),
		AccessToken:  c.String(internal.FlagAPIAccessToken),
		RefreshToken: c.String(internal.FlagAPIRefreshToken),
		Version:      c.String(internal.FlagAPIVersion),
		LogLevel:     c.String(internal.FlagLogLevel),
		Output:       c.String(internal.FlagOutput),
		Org:          c.String(internal.FlagOrg),
		Repo:         c.String(internal.FlagRepo),
		Engine:       c.String(internal.FlagSecretEngine),
		Type:         c.String(internal.FlagSecretType),
		GitHub: &config.GitHub{
			Token: c.String(internal.FlagCompilerGitHubToken),
			URL:   c.String(internal.FlagCompilerGitHubURL),
		},
	}

	// validate config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Validate
	err = conf.Validate()
	if err != nil {
		return err
	}

	// execute the generate call for the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Generate
	err = conf.Generate()
	if err != nil {
		return err
	}

	logrus.Info("configuration successfully created - enjoy")

	return nil
}
