// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"fmt"

	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal"

	"github.com/urfave/cli/v2"
)

// CommandGenerate defines the command for producing the config file.
var CommandGenerate = &cli.Command{
	Name:        "config",
	Description: "Use this command to generate the config file.",
	Usage:       "Generate the config file used in the CLI",
	Action:      generate,
	Flags: []cli.Flag{

		// API Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ADDR", "CONFIG_ADDR"},
			Name:    internal.FlagAPIAddress,
			Aliases: []string{"a"},
			Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
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
			Usage:   "token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_API_VERSION", "CONFIG_API_VERSION"},
			Name:    internal.FlagAPIVersion,
			Aliases: []string{"av"},
			Usage:   "API version for communication with the Vela server",
		},

		// Log Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL"},
			Name:    internal.FlagLogLevel,
			Aliases: []string{"l"},
			Usage:   "set the level of logging - options: (trace|debug|info|warn|error|fatal|panic)",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "CONFIG_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, or yaml format",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "CONFIG_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "CONFIG_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the CLI",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "CONFIG_ENGINE"},
			Name:    internal.FlagSecretEngine,
			Aliases: []string{"e"},
			Usage:   "provide the secret engine for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "CONFIG_TYPE"},
			Name:    internal.FlagSecretType,
			Aliases: []string{"ty"},
			Usage:   "provide the secret type for the CLI",
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
  1. Generate the config file with a Vela server address.
    $ {{.HelpName}} --api.addr https://vela.example.com
  2. Generate the config file with Vela server token.
    $ {{.HelpName}} --api.token fakeToken
  3. Generate the config file with secret engine and type.
    $ {{.HelpName}} --secret.engine native --secret.type org
  4. Generate the config file with trace level logging.
    $ {{.HelpName}} --log.level trace
  5. Generate the config file when environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/config/generate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to produce the
// config file.
func generate(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action:       internal.ActionGenerate,
		File:         c.String(internal.FlagConfig),
		Addr:         c.String(internal.FlagAPIAddress),
		Token:        c.String(internal.FlagAPIToken),
		AccessToken:  c.String(internal.FlagAPIAccessToken),
		RefreshToken: c.String(internal.FlagAPIRefreshToken),
		Version:      c.String(internal.FlagAPIVersion),
		LogLevel:     c.String(internal.FlagLogLevel),
		NoGit:        c.String(internal.FlagNoGit),
		Output:       c.String(internal.FlagOutput),
		Org:          c.String(internal.FlagOrg),
		Repo:         c.String(internal.FlagRepo),
		Engine:       c.String(internal.FlagSecretEngine),
		Type:         c.String(internal.FlagSecretType),
		GitHub: &config.GitHub{
			Token: c.String(internal.FlagCompilerGitHubToken),
			URL:   c.String(internal.FlagCompilerGitHubURL),
		},
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
