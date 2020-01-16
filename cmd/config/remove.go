// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// RemoveCmd defines the command for deleting a configuration file.
var RemoveCmd = cli.Command{
	Name:        "config",
	Description: "Use this command to remove a field or all fields in the config file.",
	Usage:       "Remove a field or all fields in the config file.",
	Action:      remove,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "addr",
			Usage: "removes the addr field from the config",
		},
		cli.BoolFlag{
			Name:  "token",
			Usage: "removes the token field from the config",
		},
		cli.BoolFlag{
			Name:  "api-version",
			Usage: "removes the api-version field from the config",
		},
		cli.BoolFlag{
			Name:  "log-level",
			Usage: "removes the log-level field from the config",
		},
		cli.BoolFlag{
			Name:  "org",
			Usage: "removes the org field from the config",
		},
		cli.BoolFlag{
			Name:  "repo",
			Usage: "removes the repo field from the config",
		},
		cli.BoolFlag{
			Name:  "secret-engine",
			Usage: "removes the secret-engine field from the config",
		},
		cli.BoolFlag{
			Name:  "secret-type",
			Usage: "removes the secret-type field from the config",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Remove the CLI config file.
    $ {{.HelpName}}
 2. Remove the address from the CLI config file.
    $ {{.HelpName}} --address
 3. Remove the API version from the CLI config file.
    $ {{.HelpName}} --api-version
 4. Remove the log level from the CLI config file.
    $ {{.HelpName}} --log-level
`, cli.CommandHelpTemplate),
}

// helper function to execute a remove repo cli command
func remove(c *cli.Context) error {
	file := c.GlobalString("config")

	_, err := os.Stat(file)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("unable to search for config file @ %s: %v", file, err)
		}

		return fmt.Errorf("unable to find config file @ %s", file)
	}

	isFlag := false
	conf := &config{
		Addr:         c.GlobalString("addr"),
		Token:        c.GlobalString("token"),
		Version:      c.GlobalString("api-version"),
		LogLevel:     c.GlobalString("log-level"),
		Org:          c.GlobalString("org"),
		Repo:         c.GlobalString("repo"),
		SecretEngine: c.GlobalString("secret-engine"),
		SecretType:   c.GlobalString("secret-type"),
	}

	// only remove global variables if flags are provided
	if c.Bool("addr") {
		conf.Addr = ""
		isFlag = true
	}

	if c.Bool("token") {
		conf.Token = ""
		isFlag = true
	}

	if c.Bool("api-version") {
		conf.Version = ""
		isFlag = true
	}

	if c.Bool("log-level") {
		conf.LogLevel = ""
		isFlag = true
	}

	if c.Bool("org") {
		conf.Org = ""
		isFlag = true
	}

	if c.Bool("repo") {
		conf.Repo = ""
		isFlag = true
	}

	if c.Bool("secret-engine") {
		conf.SecretEngine = ""
		isFlag = true
	}

	if c.Bool("secret-type") {
		conf.SecretType = ""
		isFlag = true
	}

	if !isFlag {
		err = os.Remove(file)
		if err != nil {
			return fmt.Errorf("unable to remove config file @ %s", file)
		}

		return nil
	}

	data, err := yaml.Marshal(&conf)
	if err != nil {
		return fmt.Errorf("unable to update config content: %v", err)
	}

	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("unable to create yaml config file @ %s: %v", file, err)
	}

	return nil
}
