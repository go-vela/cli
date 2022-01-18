// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"fmt"
	"io"

	"github.com/manifoldco/promptui"

	"github.com/sirupsen/logrus"
)

// PromptBrowserConfirm provides a prompt to confirm opening a browser window.
func (c *Config) PromptBrowserConfirm(in io.ReadCloser) error {
	logrus.Debug("executing prompt to confirm opening a browser")

	txtPressEnter := promptui.Styler(promptui.FGBold)("Press Enter")

	p := promptui.Prompt{
		Label: fmt.Sprintf("Open %s in your browser and complete "+
			"authentication (%s to confirm)", c.Address, txtPressEnter),
		IsConfirm: true,
		Default:   "y",
		Stdin:     in,
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
		Label: "Authentication complete. Continue to save configuration " +
			"(existing config will be overwritten)",
		IsConfirm: true,
		Default:   "n",
		Stdin:     in,
	}

	_, err := p.Run()
	if err != nil {
		return err
	}

	return nil
}
