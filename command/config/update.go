// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal"
)

// CommandUpdate defines the command for modifying one or more fields from the config file.
var CommandUpdate = &cli.Command{
	Name:        "config",
	Description: "Use this command to update one or more fields from the config file.",
	Usage:       "Update the config file used in the CLI",
	Action:      update,
	Flags: []cli.Flag{

		// API Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ADDR", "CONFIG_ADDR"),
			Name:    internal.FlagAPIAddress,
			Aliases: []string{"a"},
			Usage:   "update the API addr in the config file",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ACCESS_TOKEN", "CONFIG_ACCESS_TOKEN"),
			Name:    internal.FlagAPIAccessToken,
			Aliases: []string{"at"},
			Usage:   "access token used for communication with the Vela server",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REFRESH_TOKEN", "CONFIG_REFRESH_TOKEN"),
			Name:    internal.FlagAPIRefreshToken,
			Aliases: []string{"rt"},
			Usage:   "refresh token used for communication with the Vela server",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_TOKEN", "CONFIG_TOKEN"),
			Name:    internal.FlagAPIToken,
			Aliases: []string{"t"},
			Usage:   "update the API token in the config file",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_API_VERSION", "CONFIG_API_VERSION"),
			Name:    internal.FlagAPIVersion,
			Aliases: []string{"av"},
			Usage:   "update the API version in the config file",
		},

		// Log Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL"),
			Name:    internal.FlagLogLevel,
			Aliases: []string{"l"},
			Usage:   "update the log level in the config file",
		},

		// No Git Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_NO_GIT", "CONFIG_NO_GIT", "NO_GIT"),
			Name:    internal.FlagNoGit,
			Aliases: []string{"ng"},
			Usage:   "update the status of syncing git repo and org with .git/ directory",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "CONFIG_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "update the output in the config file",
		},

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "CONFIG_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "update the org in the config file",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "CONFIG_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "update the repo in the config file",
		},

		// Secret Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ENGINE", "CONFIG_ENGINE"),
			Name:    internal.FlagSecretEngine,
			Aliases: []string{"e"},
			Usage:   "update the secret engine in the config file",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_TYPE", "CONFIG_TYPE"),
			Name:    internal.FlagSecretType,
			Aliases: []string{"ty"},
			Usage:   "update the secret type in the config file",
		},

		// Compiler Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_COMPILER_GITHUB_TOKEN", "COMPILER_GITHUB_TOKEN"),
			Name:    internal.FlagCompilerGitHubToken,
			Aliases: []string{"ct"},
			Usage:   "github compiler token",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_COMPILER_GITHUB_URL", "COMPILER_GITHUB_URL"),
			Name:    internal.FlagCompilerGitHubURL,
			Aliases: []string{"cgu"},
			Usage:   "github url, used by compiler, for pulling registry templates",
		},

		// Test Flags (Hidden)

		&cli.BoolFlag{
			Hidden: true,
			Name:   "fs.mem-map",
			Usage:  "use memory mapped files for the config file (for testing)",
			Value:  false,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update the addr field in the config file.
    $ {{.FullName}} --api.addr https://vela.example.com
  2. Update the token field in the config file.
    $ {{.FullName}} --api.token fakeToken
  3. Update the secret engine and type fields in the config file.
    $ {{.FullName}} --secret.engine native --secret.type org
  4. Update the log level field in the config file.
    $ {{.FullName}} --log.level trace
  5. Update the no git field in the config file.
    $ {{.FullName}} --no-git bool
  6. Update the config file when environment variables are set.
    $ {{.FullName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/config/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to modify the
// config file.
func update(_ context.Context, c *cli.Command) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action:      internal.ActionUpdate,
		File:        c.String(internal.FlagConfig),
		UpdateFlags: make(map[string]string),
		UseMemMap:   c.Bool("fs.mem-map"),
	}

	// create variables from flags provided
	addr := c.String(internal.FlagAPIAddress)
	token := c.String(internal.FlagAPIToken)
	accessToken := c.String(internal.FlagAPIAccessToken)
	refreshToken := c.String(internal.FlagAPIRefreshToken)
	version := c.String(internal.FlagAPIVersion)
	level := c.String(internal.FlagLogLevel)
	noGit := c.String(internal.FlagNoGit)
	output := c.String(internal.FlagOutput)
	colorFmt := c.String(internal.FlagColorFormat)
	colorTheme := c.String(internal.FlagColorTheme)
	org := c.String(internal.FlagOrg)
	repo := c.String(internal.FlagRepo)
	engine := c.String(internal.FlagSecretEngine)
	typee := c.String(internal.FlagSecretType)
	githubToken := c.String(internal.FlagCompilerGitHubToken)
	githubURL := c.String(internal.FlagCompilerGitHubURL)

	// check if the API addr flag should be modified
	if len(addr) > 0 {
		conf.UpdateFlags[internal.FlagAPIAddress] = addr
	}

	// check if the API token flag should be modified
	if len(token) > 0 {
		conf.UpdateFlags[internal.FlagAPIToken] = token
	}

	// check if the API access token flag should be modified
	if len(accessToken) > 0 {
		conf.UpdateFlags[internal.FlagAPIAccessToken] = accessToken
	}

	// check if the API refresh token flag should be modified
	if len(refreshToken) > 0 {
		conf.UpdateFlags[internal.FlagAPIRefreshToken] = refreshToken
	}

	// check if the API version flag should be modified
	if len(version) > 0 {
		conf.UpdateFlags[internal.FlagAPIVersion] = version
	}

	// check if the log level flag should be modified
	if len(level) > 0 {
		conf.UpdateFlags[internal.FlagLogLevel] = level
	}

	// check if the no git flag should be modified
	if len(noGit) > 0 {
		conf.UpdateFlags[internal.FlagNoGit] = noGit
	}

	// check if the output flag should be modified
	if len(output) > 0 {
		conf.UpdateFlags[internal.FlagOutput] = output
	}

	// check if the color flag should be modified
	if c.IsSet(internal.FlagColor) {
		color := "true"
		if !internal.StringToBool(c.String(internal.FlagColor)) {
			color = "false"
		}

		conf.UpdateFlags[internal.FlagColor] = color
	}

	// check if the color format flag should be modified
	if len(colorFmt) > 0 {
		conf.UpdateFlags[internal.FlagColorFormat] = colorFmt
	}

	// check if the color theme flag should be modified
	if len(colorTheme) > 0 {
		conf.UpdateFlags[internal.FlagColorTheme] = colorTheme
	}

	// check if the org flag should be modified
	if len(org) > 0 {
		conf.UpdateFlags[internal.FlagOrg] = org
	}

	// check if the repo flag should be modified
	if len(repo) > 0 {
		conf.UpdateFlags[internal.FlagRepo] = repo
	}

	// check if the secret engine flag should be modified
	if len(engine) > 0 {
		conf.UpdateFlags[internal.FlagSecretEngine] = engine
	}

	// check if the secret type flag should be modified
	if len(typee) > 0 {
		conf.UpdateFlags[internal.FlagSecretType] = typee
	}

	// check if the compiler github token flag should be modified
	if len(githubToken) > 0 {
		conf.UpdateFlags[internal.FlagCompilerGitHubToken] = githubToken
	}

	// check if the compiler github url flag should be modified
	if len(githubURL) > 0 {
		conf.UpdateFlags[internal.FlagCompilerGitHubURL] = githubURL
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
