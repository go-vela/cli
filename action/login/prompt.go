// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"fmt"
	"io"

	"github.com/manifoldco/promptui"

	"github.com/sirupsen/logrus"
)

// PromptUsername asks the user to provide a username via
// terminal input based off the provided configuration.
func (c *Config) PromptUsername(in io.ReadCloser) error {
	logrus.Debug("executing prompt for username for login configuration")

	// create variable to store errors
	var err error

	// create the prompt for a username
	p := promptui.Prompt{
		Label: "Please enter a username: ",
		Stdin: in,
	}

	// run the prompt to capture the username from the input
	c.Username, err = p.Run()
	if err != nil {
		return err
	}

	logrus.Trace("checking username input provided")

	// check if login username is set
	if len(c.Username) == 0 {
		return fmt.Errorf("no login username provided")
	}

	return nil
}

// PromptPassword asks the user to provide a password via
// terminal input based off the provided configuration.
func (c *Config) PromptPassword(in io.ReadCloser) error {
	logrus.Debug("executing prompt for password for login configuration")

	// create variable to store errors
	var err error

	// create the prompt for a password
	p := promptui.Prompt{
		Label: "Please enter a password: ",
		Mask:  '*',
		Stdin: in,
	}

	// run the prompt to capture the password from the input
	c.Password, err = p.Run()
	if err != nil {
		return err
	}

	logrus.Trace("checking password input provided")

	// check if login password is set
	if len(c.Password) == 0 {
		return fmt.Errorf("no login password provided")
	}

	return nil
}

// PromptOTP asks the user to provide a OTP via
// terminal input based off the provided configuration.
func (c *Config) PromptOTP(in io.ReadCloser) error {
	logrus.Debug("executing prompt for one time password (OTP) for login configuration")

	// create variable to store errors
	var err error

	// create the prompt for a OTP
	p := promptui.Prompt{
		Label: "Please enter a one time password (OTP): ",
		Stdin: in,
	}

	// run the prompt to capture the OTP from the input
	c.OTP, err = p.Run()
	if err != nil {
		return err
	}

	logrus.Trace("checking one time password (OTP) input provided")

	// check if login OTP is set
	if len(c.OTP) == 0 {
		return fmt.Errorf("no login one time password (OTP) provided")
	}

	return nil
}
