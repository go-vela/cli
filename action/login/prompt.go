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

// PromptAddress
func (c *Config) PromptAddress(in io.ReadCloser) error {
	logrus.Debug("executing prompt for address for login configuration")

	// create variable to store errors
	var err error

	// create the prompt for the address
	p := promptui.Prompt{
		Label: "Please enter the Vela server address: ",
		Stdin: in,
	}

	// run the prompt to capture the address from the input
	c.Address, err = p.Run()
	if err != nil {
		return err
	}

	logrus.Trace("checking address input provided")

	// check if address is set
	if len(c.Address) == 0 {
		return fmt.Errorf("no address provided")
	}

	return nil
}

// PromptBrowserConfirm provides a prompt to confirm opening a browser window.
func (c *Config) PromptBrowserConfirm(in io.ReadCloser, site string) error {
	logrus.Debug("executing prompt to confirm opening a browser")

	txtPressEnter := promptui.Styler(promptui.FGBold)("Press Enter")

	p := promptui.Prompt{
		Label: fmt.Sprintf("Open %s in your browser and complete authentication (%s to confirm)",
			site, txtPressEnter),
		IsConfirm: true,
		Default:   "y",
	}

	_, err := p.Run()
	if err != nil {
		return err
	}

	return nil
}

// PromptConfigConfirm provides a prompt to confirm config generation.
func (c *Config) PromptConfigConfirm(in io.ReadCloser) error {
	logrus.Debug("executing prompt to confirm config write")

	p := promptui.Prompt{
		Label:     "Ready to write config. Any existing config will be overwritten. Continue",
		IsConfirm: true,
		Default:   "n",
	}

	_, err := p.Run()
	if err != nil {
		return err
	}

	return nil
}
