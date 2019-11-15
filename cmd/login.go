// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"errors"
	"fmt"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/library"

	"github.com/manifoldco/promptui"

	"github.com/urfave/cli"
)

// loginCmd defines the command for authenticating and logging in to Vela.
var loginCmd = cli.Command{
	Name:        "login",
	Category:    "Authentication",
	Aliases:     []string{"l"},
	Description: "Use this command to authenticate and login to Vela.",
	Usage:       "Authenticate and generate a token with Vela",
	Action:      authenticate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		cli.StringFlag{
			Name:  "addr",
			Usage: "location of vela server",
		},

		// optional flags that can be supplied to a command
		cli.StringFlag{
			Name:  "username,u",
			Usage: "Override username prompt",
		},
		cli.StringFlag{
			Name:  "password,p",
			Usage: "Override username prompt",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Login to Vela with terminal prompts.
    $ {{.HelpName}} --addr https://vela.example.com
 2. Login to Vela with username and password.
    $ {{.HelpName}} --username foo --password bar
`, cli.CommandHelpTemplate),
}

func authenticate(c *cli.Context) error {
	var (
		message string
		token   string
	)

	url := c.String("addr")
	if len(url) == 0 { // use global variable if flags isn't provided
		url = c.GlobalString("addr")
	}

	// create a carval client
	client, err := vela.NewClient(url, nil)
	if err != nil {
		return err
	}

	// get username from user input
	username, err := getUsername(c)
	if err != nil {
		return err
	}

	// get password from user input
	password, err := getPassword(c)
	if err != nil {
		return err
	}

	req := library.Login{
		Username: &username,
		Password: &password,
		OTP:      vela.String(""),
	}

	auth, resp, err := client.Authorization.Login(&req)

	// If user hits an endpoint other than the
	// Vela server that can't process request
	// bomb out and throw error
	if 401 < resp.StatusCode {
		return fmt.Errorf("unable to process request")
	}
	if resp.StatusCode != 401 && err != nil {
		return err
	}

	// retry authentication in case user requires an OTP code
	switch resp.StatusCode {
	case 401:
		// get otp from user input
		otp, err := getOTP()
		if err != nil {
			return err
		}

		req = library.Login{
			Username: &username,
			Password: &password,
			OTP:      &otp,
		}

		auth, _, err := client.Authorization.Login(&req)
		if err != nil {
			return err
		}

		// craft response to user
		message = fmt.Sprintf("Generated token: %s", auth.GetToken())
		token = auth.GetToken()
	default:

		// craft response to user
		message = fmt.Sprintf("Generated token: %s", auth.GetToken())
		token = auth.GetToken()
	}

	if token != "" {
		fmt.Println(message)
	}

	return nil
}

// helper function get a username from user input
func getUsername(c *cli.Context) (string, error) {

	if len(c.String("username")) != 0 {
		return c.String("username"), nil
	}

	validate := func(input string) error {
		if len(input) < 0 {
			return errors.New("Username must not be blank")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter a username: ",
		Validate: validate,
	}

	u, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return u, nil
}

// helper function get a password from user input
func getPassword(c *cli.Context) (string, error) {

	if len(c.String("password")) != 0 {
		return c.String("password"), nil
	}

	validate := func(input string) error {
		if len(input) < 0 {
			return errors.New("Username must not be blank")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter a password: ",
		Validate: validate,
		Mask:     '*',
	}

	p, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return p, nil
}

// helper function get a one time passsword from user input
func getOTP() (string, error) {

	validate := func(input string) error {
		if len(input) < 0 {
			return errors.New("OTP must not be blank")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter a 2fa code: ",
		Validate: validate,
	}

	otp, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return otp, nil
}
