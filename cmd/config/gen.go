// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// config represents the values stored in a configuration file.
type config struct {
	Addr         string `yaml:"addr,omitempty"`
	Token        string `yaml:"token,omitempty"`
	Version      string `yaml:"api-version,omitempty"`
	LogLevel     string `yaml:"log-level,omitempty"`
	Output       string `yaml:"output,omitempty"`
	Org          string `yaml:"org,omitempty"`
	Repo         string `yaml:"repo,omitempty"`
	SecretEngine string `yaml:"secret-engine,omitempty"`
	SecretType   string `yaml:"secret-type,omitempty"`
}

// GenCmd defines the command for creating a configuration file.
var GenCmd = cli.Command{
	Name:        "config",
	Description: "Use this command to add a config file.",
	Usage:       "Generate a config yaml in a directory",
	Action:      gen,
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
 1. Generate CLI config from environment.
    $ {{.HelpName}}
 2. Generate CLI config with address and token.
    $ {{.HelpName}} --addr https://vela.example.com --token fakeToken
 3. Generate CLI config with specific API version.
    $ {{.HelpName}} --api-version v1
 4. Generate CLI config with debug log level.
    $ {{.HelpName}} --log-level debug
`, cli.CommandHelpTemplate),
}

// helper function to add a configuration file
func gen(c *cli.Context) error {
	conf := &config{
		Addr:         c.String("addr"),
		Token:        c.String("token"),
		Version:      c.String("api-version"),
		LogLevel:     c.String("log-level"),
		Org:          c.String("org"),
		Repo:         c.String("repo"),
		SecretEngine: c.String("secret-engine"),
		SecretType:   c.String("secret-type"),
	}

	// use global variables if flags aren't provided
	if len(conf.Addr) == 0 {
		conf.Addr = c.GlobalString("addr")
	}

	if len(conf.Token) == 0 {
		conf.Token = c.GlobalString("token")
	}

	if len(conf.Version) == 0 {
		conf.Version = c.GlobalString("api-version")
	}

	if len(conf.LogLevel) == 0 {
		conf.LogLevel = c.GlobalString("log-level")
	}

	if len(conf.Org) == 0 {
		conf.Org = c.GlobalString("org")
	}

	if len(conf.Repo) == 0 {
		conf.Repo = c.GlobalString("repo")
	}

	if len(conf.SecretEngine) == 0 {
		conf.SecretEngine = c.GlobalString("secret-engine")
	}

	if len(conf.SecretType) == 0 {
		conf.SecretType = c.GlobalString("secret-type")
	}

	data, err := yaml.Marshal(&conf)
	if err != nil {
		return fmt.Errorf("unable to create config content: %v", err)
	}

	file := c.GlobalString("config")
	directory := filepath.Dir(file)

	err = os.MkdirAll(directory, 0777)
	if err != nil {
		return fmt.Errorf("unable to create directory path to config @ %s: %v", directory, err)
	}

	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("unable to create yaml config file @ %s: %v", file, err)
	}

	fmt.Printf("Yaml config file created @ %s\n", file)

	return nil
}
