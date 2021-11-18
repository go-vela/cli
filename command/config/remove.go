// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"

	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal"

	"github.com/urfave/cli/v2"
)

// CommandRemove defines the command for deleting one or more fields from the config file.
var CommandRemove = &cli.Command{
	Name:        "config",
	Description: "Use this command to remove one or more fields from the config file.",
	Usage:       "Remove the config file used in the CLI",
	Action:      remove,
	Flags: []cli.Flag{

		// API Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ADDR", "CONFIG_ADDR"},
			Name:    internal.FlagAPIAddress,
			Aliases: []string{"a"},
			Usage:   "removes the API addr from the config file",
			Value:   "false",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ACCESS_TOKEN", "CONFIG_ACCESS_TOKEN"},
			Name:    internal.FlagAPIAccessToken,
			Aliases: []string{"at"},
			Usage:   "access token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REFRESH_TOKEN", "CONFIG_REFRESH_TOKEN"},
			Name:    internal.FlagAPIRefreshToken,
			Aliases: []string{"rt"},
			Usage:   "refresh token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN"},
			Name:    internal.FlagAPIToken,
			Aliases: []string{"t"},
			Usage:   "removes the API token from the config file",
			Value:   "false",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_API_VERSION", "CONFIG_API_VERSION"},
			Name:    internal.FlagAPIVersion,
			Aliases: []string{"av"},
			Usage:   "removes the API version from the config file",
			Value:   "false",
		},

		// Log Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL"},
			Name:    internal.FlagLogLevel,
			Aliases: []string{"l"},
			Usage:   "removes the log level from the config file",
			Value:   "false",
		},

		// Git Sync Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_GIT_SYNC", "CONFIG_GIT_SYNC", "GIT_SYNC"},
			Name:    internal.FlagNoGit,
			Aliases: []string{"gs"},
			Usage:   "removes the status of syncing git repo and org with .git/ directory",
			Value:   "false",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "CONFIG_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "removes the output from the config file",
			Value:   "false",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "CONFIG_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "removes the org from the config file",
			Value:   "false",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "CONFIG_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "removes the repo from the config file",
			Value:   "false",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "CONFIG_ENGINE"},
			Name:    internal.FlagSecretEngine,
			Aliases: []string{"e"},
			Usage:   "removes the secret engine from the config file",
			Value:   "false",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "CONFIG_TYPE"},
			Name:    internal.FlagSecretType,
			Aliases: []string{"ty"},
			Usage:   "removes the secret type from the config file",
			Value:   "false",
		},

		// Compiler Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_COMPILER_GITHUB_TOKEN", "COMPILER_GITHUB_TOKEN"},
			Name:    internal.FlagCompilerGitHubToken,
			Aliases: []string{"ct"},
			Usage:   "github compiler token",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_COMPILER_GITHUB_URL", "COMPILER_GITHUB_URL"},
			Name:    internal.FlagCompilerGitHubURL,
			Aliases: []string{"cgu"},
			Usage:   "github url, used by compiler, for pulling registry templates",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Remove the config file.
    $ {{.HelpName}}
  2. Remove the addr field from the config file.
    $ {{.HelpName}} --api.addr true
  3. Remove the token field from the config file.
    $ {{.HelpName}} --api.token true
  4. Remove the secret engine and type fields from the config file.
    $ {{.HelpName}} --secret.engine true --secret.type true
  5. Remove the log level field from the config file.
    $ {{.HelpName}} --log.level true

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/config/remove/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to delete one or
// more fields from the config file.
func remove(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action: internal.ActionRemove,
		File:   c.String(internal.FlagConfig),
	}

	// check if the API addr flag should be removed
	if c.Bool(internal.FlagAPIAddress) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagAPIAddress)
	}

	// check if the API token flag should be removed
	if c.Bool(internal.FlagAPIToken) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagAPIToken)
	}

	// check if the API access token flag should be removed
	if c.Bool(internal.FlagAPIAccessToken) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagAPIAccessToken)
	}

	// check if the API refresh token flag should be removed
	if c.Bool(internal.FlagAPIRefreshToken) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagAPIRefreshToken)
	}

	// check if the API version flag should be removed
	if c.Bool(internal.FlagAPIVersion) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagAPIVersion)
	}

	// check if the log level flag should be removed
	if c.Bool(internal.FlagLogLevel) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagLogLevel)
	}

	// check if the git sync flag should be removed
	if c.Bool(internal.FlagNoGit) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagNoGit)
	}

	// check if the output flag should be removed
	if c.Bool(internal.FlagOutput) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagOutput)
	}

	// check if the org flag should be removed
	if c.Bool(internal.FlagOrg) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagOrg)
	}

	// check if the repo flag should be removed
	if c.Bool(internal.FlagRepo) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagRepo)
	}

	// check if the engine flag should be removed
	if c.Bool(internal.FlagSecretEngine) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagSecretEngine)
	}

	// check if the type flag should be removed
	if c.Bool(internal.FlagSecretType) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagSecretType)
	}

	// check if the compiler github token flag should be removed
	if c.Bool(internal.FlagCompilerGitHubToken) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagCompilerGitHubToken)
	}

	// check if the compiler github url flag should be removed
	if c.Bool(internal.FlagCompilerGitHubURL) {
		conf.RemoveFlags = append(conf.RemoveFlags, internal.FlagCompilerGitHubURL)
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
