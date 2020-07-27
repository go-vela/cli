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
			EnvVars: []string{"VELA_USERNAME", "LOGIN_USERNAME"},
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "overrides the prompt for a username",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_PASSWORD", "LOGIN_PASSWORD"},
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "overrides the prompt for a password",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_OTP", "LOGIN_OTP"},
			Name:    "otp",
			Aliases: []string{"o"},
			Usage:   "overrides the prompt for a OTP",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Login to Vela with terminal prompts.
    $ {{.HelpName}} --api.addr https://vela.example.com
  2. Login to Vela with no prompts for username and password
    $ {{.HelpName}} --username foo --password bar

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
	l := &login.Config{
		Action:   loginAction,
		Username: c.String("username"),
		Password: c.String("password"),
		OTP:      c.String("otp"),
	}

	// check if username was provided from flags
	if len(l.Username) == 0 {
		// prompt user to provide username via terminal input
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.PromptUsername
		err = l.PromptUsername(os.Stdin)
		if err != nil {
			return err
		}
	}

	// check if password was provided from flags
	if len(l.Password) == 0 {
		// prompt user to provide password via terminal input
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.PromptPassword
		err = l.PromptPassword(os.Stdin)
		if err != nil {
			return err
		}
	}

	// validate login configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.Validate
	err = l.Validate()
	if err != nil {
		return err
	}

	// execute the login call for the login configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.Login
	err = l.Login(client)
	if err == nil {
		// return instantly if no error occurred during the login
		return nil
	}

	// check if the retry mechanism is set
	//
	// an error might be returned with the retry mechanism set if
	// the source system Vela is integrated with requires a OTP
	if !l.Retry {
		return err
	}

	// prompt user to provide OTP via terminal input
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.PromptOTP
	err = l.PromptOTP(os.Stdin)
	if err != nil {
		return err
	}

	// validate login configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.Validate
	err = l.Validate()
	if err != nil {
		return err
	}

	// execute the login call for the login configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/login?tab=doc#Config.Login
	return l.Login(client)
}
