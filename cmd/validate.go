// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/go-vela/lexi/compiler/native"
	yLib "github.com/go-vela/types/yaml"

	"github.com/urfave/cli"
)

// validateCmd defines the command that ensures you have a runnable pipeline
var validateCmd = cli.Command{
	Name:        "validate",
	Category:    "Pipeline Management",
	Aliases:     []string{"va"},
	Description: "Use this command to validate your pipeline",
	Usage:       "Validate the vela config file in the current directory",
	Action:      validate,
	Flags:       []cli.Flag{},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Validate the vela config in current directory.
		$ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to run the validate pipeline command
func validate(c *cli.Context) error {
	client, err := native.New(c)
	if err != nil {
		return err
	}

	data, err := openFile()
	if err != nil {
		return err
	}
	var b yLib.Build

	err = yaml.Unmarshal(data, &b)
	if err != nil {
		return err
	}

	err = client.Validate(&b)
	if err != nil {
		return err
	}

	if err == nil {
		fmt.Println("Config is valid.")
	}

	return nil
}

// helper function to open the vela pipeline and get contents
func openFile() ([]byte, error) {
	d, err := ioutil.ReadFile(".vela.yml")
	if err != nil {
		d, err = ioutil.ReadFile(".vela.yaml")
		if err != nil {
			return nil, fmt.Errorf("No vela file in this directory")
		}
	}
	return d, nil
}
