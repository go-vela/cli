// SPDX-License-Identifier: Apache-2.0

package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v3"

	"github.com/go-vela/cli/internal"
)

// Load reads the config file and sets the values based off the provided configuration.
//
//nolint:funlen,gocyclo // ignore cyclomatic complexity and function length
func (c *Config) Load(ctx *cli.Context) error {
	logrus.Debug("executing load for config file configuration")

	// check if we're operating on the config resource
	for _, arg := range ctx.Args().Slice() {
		// check if we're operating on the config resource
		if strings.EqualFold(arg, "config") {
			logrus.Debugf("config arg provided in %v - skipping load for config file %s", ctx.Args().Slice(), c.File)

			return nil
		}
	}

	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	// send Filesystem call to read config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.ReadFile
	data, err := a.ReadFile(c.File)
	if err != nil {
		return err
	}

	// create the config file object
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#ConfigFile
	config := new(ConfigFile)

	// update the config object with the current content
	//
	// https://pkg.go.dev/gopkg.in/yaml.v3?tab=doc#Unmarshal
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}

	// check if the config file is empty
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#ConfigFile.Empty
	if config.Empty() {
		logrus.Warningf("empty/unsupported config loaded from %s - fix this with `vela generate config`", c.File)

		return nil
	}

	// capture a list of all available flags to set in the current context
	flags := []cli.Flag{}
	// check if the app is provided in the context
	if ctx.App != nil {
		// append the flags from the app provided
		flags = append(flags, ctx.App.VisibleFlags()...)
	}
	// check if the command is provided in the context
	if ctx.Command != nil {
		// append the flags from the context provided
		flags = append(flags, ctx.Command.VisibleFlags()...)
	}

	// iterate through all available flags in the current context
	for _, flag := range flags {
		// capture string value for flag
		f := strings.Join(flag.Names(), " ")

		// check if the API address flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagAPIAddress) &&
			!ctx.IsSet(internal.FlagAPIAddress) &&
			len(config.API.Address) > 0 {
			// set the API address field to value from config
			err = ctx.Set(internal.FlagAPIAddress, config.API.Address)
			if err != nil {
				return err
			}

			continue
		}

		// check if the API token flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagAPIToken) &&
			!ctx.IsSet(internal.FlagAPIToken) &&
			len(config.API.Token) > 0 {
			// set the API token field to value from config
			err = ctx.Set(internal.FlagAPIToken, config.API.Token)
			if err != nil {
				return err
			}

			continue
		}

		// check if the API access token flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagAPIAccessToken) &&
			!ctx.IsSet(internal.FlagAPIAccessToken) &&
			len(config.API.AccessToken) > 0 {
			// set the API access token field to value from config
			err = ctx.Set(internal.FlagAPIAccessToken, config.API.AccessToken)
			if err != nil {
				return err
			}

			continue
		}

		// check if the API refresh token flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagAPIRefreshToken) &&
			!ctx.IsSet(internal.FlagAPIRefreshToken) &&
			len(config.API.RefreshToken) > 0 {
			// set the API refresh token field to value from config
			err = ctx.Set(internal.FlagAPIRefreshToken, config.API.RefreshToken)
			if err != nil {
				return err
			}

			continue
		}

		// check if the API version flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagAPIVersion) &&
			!ctx.IsSet(internal.FlagAPIVersion) &&
			len(config.API.Version) > 0 {
			// set the API version field to value from config
			err = ctx.Set(internal.FlagAPIVersion, config.API.Version)
			if err != nil {
				return err
			}

			continue
		}

		// check if the log level flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagLogLevel) &&
			!ctx.IsSet(internal.FlagLogLevel) &&
			len(config.Log.Level) > 0 {
			// set the log level field to value from config
			err = ctx.Set(internal.FlagLogLevel, config.Log.Level)
			if err != nil {
				return err
			}
		}

		// check if the git sync flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagNoGit) &&
			!ctx.IsSet(internal.FlagNoGit) &&
			len(config.NoGit) > 0 {
			err = ctx.Set(internal.FlagNoGit, config.NoGit)
			if err != nil {
				return err
			}
		}

		// check if the output flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagOutput) &&
			!ctx.IsSet(internal.FlagOutput) &&
			len(config.Output) > 0 {
			// set the output field to value from config
			err = ctx.Set(internal.FlagOutput, config.Output)
			if err != nil {
				return err
			}

			continue
		}

		// check if the color flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagColor) &&
			!ctx.IsSet(internal.FlagColor) {
			// set the color field to value from config
			c := "true"
			if config.Color != nil && !*config.Color {
				c = "false"
			}

			err = ctx.Set(internal.FlagColor, c)
			if err != nil {
				return err
			}

			continue
		}

		// check if the color format flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagColorFormat) &&
			!ctx.IsSet(internal.FlagColorFormat) &&
			len(config.ColorFormat) > 0 {
			// set the color format field to value from config
			err = ctx.Set(internal.FlagColorFormat, config.ColorFormat)
			if err != nil {
				return err
			}

			continue
		}

		// check if the color theme flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagColorTheme) &&
			!ctx.IsSet(internal.FlagColorTheme) &&
			len(config.ColorTheme) > 0 {
			// set the color theme field to value from config
			err = ctx.Set(internal.FlagColorTheme, config.ColorTheme)
			if err != nil {
				return err
			}

			continue
		}

		// check if the org flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagOrg) &&
			!ctx.IsSet(internal.FlagOrg) &&
			len(config.Org) > 0 {
			// set the org field to value from config
			err = ctx.Set(internal.FlagOrg, config.Org)
			if err != nil {
				return err
			}

			continue
		}

		// check if the repo flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagRepo) &&
			!ctx.IsSet(internal.FlagRepo) &&
			len(config.Repo) > 0 {
			// set the repo field to value from config
			err = ctx.Set(internal.FlagRepo, config.Repo)
			if err != nil {
				return err
			}

			continue
		}

		// check if the secret engine flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagSecretEngine) &&
			!ctx.IsSet(internal.FlagSecretEngine) &&
			len(config.Secret.Engine) > 0 {
			// set the secret engine field to value from config
			err = ctx.Set(internal.FlagSecretEngine, config.Secret.Engine)
			if err != nil {
				return err
			}

			continue
		}

		// check if the secret type flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagSecretType) &&
			!ctx.IsSet(internal.FlagSecretType) &&
			len(config.Secret.Type) > 0 {
			// set the secret type field to value from config
			err = ctx.Set(internal.FlagSecretType, config.Secret.Type)
			if err != nil {
				return err
			}

			continue
		}

		// check if the compiler github token flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagCompilerGitHubToken) &&
			!ctx.IsSet(internal.FlagCompilerGitHubToken) &&
			config.Compiler != nil &&
			len(config.Compiler.GitHub.Token) > 0 {
			// set the compiler github token field to value from config
			err = ctx.Set(internal.FlagCompilerGitHubToken, config.Compiler.GitHub.Token)
			if err != nil {
				return err
			}

			continue
		}

		// check if the compiler github url flag is available
		// and if it is set in the context
		if strings.Contains(f, internal.FlagCompilerGitHubURL) &&
			!ctx.IsSet(internal.FlagCompilerGitHubURL) &&
			config.Compiler != nil &&
			len(config.Compiler.GitHub.URL) > 0 {
			// set the compiler github url field to value from config
			err = ctx.Set(internal.FlagCompilerGitHubURL, config.Compiler.GitHub.URL)
			if err != nil {
				return err
			}

			continue
		}
	}

	return nil
}
