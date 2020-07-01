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

// ConfigGenerate defines the command for producing the config file.
var ConfigGenerate = &cli.Command{
	Name:        "config",
	Description: "Use this command to generate the config file.",
	Usage:       "Generate the config file used in the CLI",
	Action:      configGenerate,
	Flags: []cli.Flag{

		// API Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ADDR", "CONFIG_ADDR"},
			Name:    client.KeyAddress,
			Aliases: []string{"a"},
			Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN"},
			Name:    client.KeyToken,
			Aliases: []string{"t"},
			Usage:   "token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_API_VERSION", "CONFIG_API_VERSION"},
			Name:    "api.version",
			Aliases: []string{"av"},
			Usage:   "API version for communication with the Vela server",
		},

		// Log Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL"},
			Name:    "log.level",
			Aliases: []string{"l"},
			Usage:   "set the level of logging for the CLI",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "CONFIG_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "set the type of output for the CLI",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "CONFIG_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "CONFIG_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the CLI",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "CONFIG_ENGINE"},
			Name:    "engine",
			Aliases: []string{"e"},
			Usage:   "provide the secret engine for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "CONFIG_TYPE"},
			Name:    "type",
			Aliases: []string{"ty"},
			Usage:   "provide the secret type for the CLI",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Generate the config file with a Vela server address.
    $ {{.HelpName}} --addr https://vela.example.com
  2. Generate the config file with Vela server token.
    $ {{.HelpName}} --token fakeToken
  3. Generate the config file with secret engine and type.
    $ {{.HelpName}} --engine native --type org
  4. Generate the config file with trace level logging.
    $ {{.HelpName}} --log.level trace
  5. Generate the config file when environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/config/generate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// produce the config file.
func configGenerate(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action:   generateAction,
		File:     c.String("file"),
		Addr:     c.String(client.KeyAddress),
		Token:    c.String(client.KeyToken),
		Version:  c.String("api.version"),
		LogLevel: c.String("log.level"),
		Output:   c.String("output"),
		Org:      c.String("org"),
		Repo:     c.String("repo"),
		Engine:   c.String("engine"),
		Type:     c.String("type"),
	}

	// validate config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Validate
	err := conf.Validate()
	if err != nil {
		return err
	}

	// execute the generate call for the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Generate
	return conf.Generate()
}
