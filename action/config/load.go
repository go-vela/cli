// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"github.com/go-vela/cli/internal"

	"github.com/sirupsen/logrus"

	"github.com/spf13/afero"

	"github.com/urfave/cli/v2"

	yaml "gopkg.in/yaml.v2"
)

// Load reads the config file and sets the values based off the provided configuration.
func (c *Config) Load(ctx *cli.Context) error {
	logrus.Debug("executing load for config file configuration")

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
	// https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc#Unmarshal
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}

	// check if the config file is empty
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#ConfigFile.Empty
	if config.Empty() {
		logrus.Debugf("empty config loaded from %s", c.File)

		return nil
	}

	// check if the API address is set in the context
	if !ctx.IsSet(internal.FlagAPIAddress) && len(config.API.Address) > 0 {
		// set the API address field to value from config
		err = ctx.Set(internal.FlagAPIAddress, config.API.Address)
		if err != nil {
			return err
		}
	}

	// check if the API token is set in the context
	if !ctx.IsSet(internal.FlagAPIToken) && len(config.API.Token) > 0 {
		// set the API token field to value from config
		err = ctx.Set(internal.FlagAPIToken, config.API.Token)
		if err != nil {
			return err
		}
	}

	// check if the API version is set in the context
	if !ctx.IsSet(internal.FlagAPIVersion) && len(config.API.Version) > 0 {
		// set the API version field to value from config
		err = ctx.Set(internal.FlagAPIVersion, config.API.Version)
		if err != nil {
			return err
		}
	}

	// check if the log level is set in the context
	if !ctx.IsSet(internal.FlagLogLevel) && len(config.Log.Level) > 0 {
		// set the log level field to value from config
		err = ctx.Set(internal.FlagLogLevel, config.Log.Level)
		if err != nil {
			return err
		}
	}

	// check if the output is set in the context
	if !ctx.IsSet(internal.FlagOutput) && len(config.Output) > 0 {
		// set the output field to value from config
		err = ctx.Set(internal.FlagOutput, config.Output)
		if err != nil {
			return err
		}
	}

	// check if the org is set in the context
	if !ctx.IsSet(internal.FlagOrg) && len(config.Org) > 0 {
		// set the org field to value from config
		err = ctx.Set(internal.FlagOrg, config.Org)
		if err != nil {
			return err
		}
	}

	// check if the repo is set in the context
	if !ctx.IsSet(internal.FlagRepo) && len(config.Repo) > 0 {
		// set the repo field to value from config
		err = ctx.Set(internal.FlagRepo, config.Repo)
		if err != nil {
			return err
		}
	}

	// check if the secret engine is set in the context
	if !ctx.IsSet(internal.FlagSecretEngine) && len(config.Secret.Engine) > 0 {
		// set the secret engine field to value from config
		err = ctx.Set(internal.FlagSecretEngine, config.Secret.Engine)
		if err != nil {
			return err
		}
	}

	// check if the secret type is set in the context
	if !ctx.IsSet(internal.FlagSecretType) && len(config.Secret.Type) > 0 {
		// set the secret type field to value from config
		err = ctx.Set(internal.FlagSecretType, config.Secret.Type)
		if err != nil {
			return err
		}
	}

	return nil
}
