// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// UpdateCmd defines the command for updating a configuration file.
var UpdateCmd = cli.Command{
	Name:        "config",
	Description: "Use this command to update a config file.",
	Usage:       "Update a config file",
	Action:      update,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Usage: "location of vela server",
		},
		cli.StringFlag{
			Name:  "token",
			Usage: "User token for Vela server",
		},
		cli.StringFlag{
			Name:  "api-version",
			Usage: "api version to use for Vela server",
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
		},
		cli.StringFlag{
			Name:  "org",
			Usage: "Provide the organization for the repository",
		},
		cli.StringFlag{
			Name:  "repo",
			Usage: "Provide the repository contained within the organization",
		},
		cli.StringFlag{
			Name:  "secret-engine",
			Usage: "Provide the engine for where the secret to be stored",
		},
		cli.StringFlag{
			Name:  "secret-type",
			Usage: "Provide the kind of secret to be stored",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Update CLI config with address.
    $ {{.HelpName}} --addr https://vela.example.com
 2. Update CLI config with personal token.
    $ {{.HelpName}} --token fakeToken
 3. Update CLI config with specific API version.
    $ {{.HelpName}} --api-version v1
 4. Update CLI config with debug log level.
    $ {{.HelpName}} --log-level debug
`, cli.CommandHelpTemplate),
}

// helper function to create a configuration file
func update(c *cli.Context) error {
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

	// only update global variables if flags are provided
	if len(c.String("addr")) > 0 {
		conf.Addr = c.String("addr")
	}
	if len(c.String("token")) > 0 {
		conf.Token = c.String("token")
	}
	if len(c.String("api-version")) > 0 {
		conf.Version = c.String("api-version")
	}
	if len(c.String("log-level")) > 0 {
		conf.LogLevel = c.String("log-level")
	}
	if len(c.String("org")) > 0 {
		conf.Org = c.String("org")
	}
	if len(c.String("repo")) > 0 {
		conf.Repo = c.String("repo")
	}
	if len(c.String("secret-engine")) > 0 {
		conf.SecretEngine = c.String("secret-engine")
	}
	if len(c.String("secret-type")) > 0 {
		conf.SecretType = c.String("secret-type")
	}

	data, err := yaml.Marshal(&conf)
	if err != nil {
		return fmt.Errorf("Unable to update config content: %v", err)
	}

	file := c.GlobalString("config")

	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("Unable to create yaml config file @ %s: %v", file, err)
	}

	return nil
}
