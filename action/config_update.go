// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// ConfigUpdate defines the command for modifying one or more fields from the config file.
var ConfigUpdate = &cli.Command{
	Name:        "config",
	Description: "Use this command to update one or more fields from the config file.",
	Usage:       "Update the config file used in the CLI",
	Action:      configUpdate,
	Flags: []cli.Flag{

		// API Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ADDR", "CONFIG_ADDR"},
			Name:    client.KeyAddress,
			Aliases: []string{"a"},
			Usage:   "update the API addr in the config file",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN"},
			Name:    client.KeyToken,
			Aliases: []string{"t"},
			Usage:   "update the API token in the config file",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_API_VERSION", "CONFIG_API_VERSION"},
			Name:    "api.version",
			Aliases: []string{"av"},
			Usage:   "update the API version in the config file",
		},

		// Log Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL"},
			Name:    "log.level",
			Aliases: []string{"l"},
			Usage:   "update the log level in the config file",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "CONFIG_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "update the output in the config file",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "CONFIG_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "update the org in the config file",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "CONFIG_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "update the repo in the config file",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "CONFIG_ENGINE"},
			Name:    "engine",
			Aliases: []string{"e"},
			Usage:   "update the secret engine in the config file",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "CONFIG_TYPE"},
			Name:    "type",
			Aliases: []string{"ty"},
			Usage:   "update the secret type in the config file",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update the addr field in the config file.
    $ {{.HelpName}} --addr https://vela.example.com
  2. Update the token field in the config file.
    $ {{.HelpName}} --token fakeToken
  3. Update the secret engine and type fields in the config file.
    $ {{.HelpName}} --engine native --type org
  4. Update the log level field in the config file.
    $ {{.HelpName}} --log.level trace
  5. Update the config file when environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/config/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// modify the config file.
func configUpdate(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action:      updateAction,
		File:        c.String("file"),
		UpdateFlags: make(map[string]string),
	}

	// create variables from flags provided
	addr := c.String(client.KeyAddress)
	token := c.String(client.KeyToken)
	version := c.String("api.version")
	level := c.String("log.level")
	output := c.String("output")
	org := c.String("org")
	repo := c.String("repo")
	engine := c.String("engine")
	typee := c.String("type")

	// check if the API addr flag should be modified
	if len(addr) > 0 {
		conf.UpdateFlags[client.KeyAddress] = addr
	}

	// check if the API token flag should be modified
	if len(token) > 0 {
		conf.UpdateFlags[client.KeyToken] = token
	}

	// check if the API version flag should be modified
	if len(version) > 0 {
		conf.UpdateFlags["api.version"] = version
	}

	// check if the log level flag should be modified
	if len(level) > 0 {
		conf.UpdateFlags["log.level"] = level
	}

	// check if the output flag should be modified
	if len(output) > 0 {
		conf.UpdateFlags["output"] = output
	}

	// check if the org flag should be modified
	if len(org) > 0 {
		conf.UpdateFlags["org"] = org
	}

	// check if the repo flag should be modified
	if len(repo) > 0 {
		conf.UpdateFlags["repo"] = repo
	}

	// check if the secret engine flag should be modified
	if len(engine) > 0 {
		conf.UpdateFlags["engine"] = engine
	}

	// check if the secret type flag should be modified
	if len(typee) > 0 {
		conf.UpdateFlags["type"] = typee
	}

	// validate config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Validate
	err := conf.Validate()
	if err != nil {
		return err
	}

	// execute the update call for the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Update
	return conf.Update()
}
