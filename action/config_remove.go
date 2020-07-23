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

// ConfigRemove defines the command for deleting one or more fields from the config file.
var ConfigRemove = &cli.Command{
	Name:        "config",
	Description: "Use this command to remove one or more fields from the config file.",
	Usage:       "Remove the config file used in the CLI",
	Action:      configRemove,
	Flags: []cli.Flag{

		// API Flags

		&cli.BoolFlag{
			EnvVars: []string{"VELA_ADDR", "CONFIG_ADDR"},
			Name:    client.KeyAddress,
			Aliases: []string{"a"},
			Usage:   "removes the API addr from the config file",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN"},
			Name:    client.KeyToken,
			Aliases: []string{"t"},
			Usage:   "removes the API token from the config file",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_API_VERSION", "CONFIG_API_VERSION"},
			Name:    "api.version",
			Aliases: []string{"av"},
			Usage:   "removes the API version from the config file",
		},

		// Log Flags

		&cli.BoolFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL"},
			Name:    "log.level",
			Aliases: []string{"l"},
			Usage:   "removes the log level from the config file",
		},

		// Output Flags

		&cli.BoolFlag{
			EnvVars: []string{"VELA_OUTPUT", "CONFIG_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "removes the output from the config file",
		},

		// Repo Flags

		&cli.BoolFlag{
			EnvVars: []string{"VELA_ORG", "CONFIG_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "removes the org from the config file",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_REPO", "CONFIG_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "removes the repo from the config file",
		},

		// Secret Flags

		&cli.BoolFlag{
			EnvVars: []string{"VELA_ENGINE", "CONFIG_ENGINE"},
			Name:    "secret.engine",
			Aliases: []string{"e"},
			Usage:   "removes the secret engine from the config file",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_TYPE", "CONFIG_TYPE"},
			Name:    "secret.type",
			Aliases: []string{"ty"},
			Usage:   "removes the secret type from the config file",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Remove the config file.
    $ {{.HelpName}}
  2. Remove the addr field from the config file.
    $ {{.HelpName}} --api.addr
  3. Remove the token field from the config file.
    $ {{.HelpName}} --api.token
  4. Remove the secret engine and type fields from the config file.
    $ {{.HelpName}} --secret.engine --secret.type
  5. Remove the log level field from the config file.
    $ {{.HelpName}} --log.level

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/config/remove/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// delete one or more fields from the
// config file.
func configRemove(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action: removeAction,
		File:   c.String("config"),
	}

	// check if the API addr flag should be removed
	if c.Bool(client.KeyAddress) {
		conf.RemoveFlags = append(conf.RemoveFlags, client.KeyAddress)
	}

	// check if the API token flag should be removed
	if c.Bool(client.KeyToken) {
		conf.RemoveFlags = append(conf.RemoveFlags, client.KeyToken)
	}

	// check if the API version flag should be removed
	if c.Bool("api.version") {
		conf.RemoveFlags = append(conf.RemoveFlags, "api.version")
	}

	// check if the log level flag should be removed
	if c.Bool("log.level") {
		conf.RemoveFlags = append(conf.RemoveFlags, "log.level")
	}

	// check if the output flag should be removed
	if c.Bool("output") {
		conf.RemoveFlags = append(conf.RemoveFlags, "output")
	}

	// check if the org flag should be removed
	if c.Bool("org") {
		conf.RemoveFlags = append(conf.RemoveFlags, "org")
	}

	// check if the repo flag should be removed
	if c.Bool("repo") {
		conf.RemoveFlags = append(conf.RemoveFlags, "repo")
	}

	// check if the engine flag should be removed
	if c.Bool("secret.engine") {
		conf.RemoveFlags = append(conf.RemoveFlags, "secret.engine")
	}

	// check if the type flag should be removed
	if c.Bool("secret.type") {
		conf.RemoveFlags = append(conf.RemoveFlags, "secret.type")
	}

	// validate config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Validate
	err := conf.Validate()
	if err != nil {
		return err
	}

	// execute the remove call for the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Remove
	return conf.Remove()
}
