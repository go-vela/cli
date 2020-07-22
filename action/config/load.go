// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"github.com/go-vela/cli/internal/client"

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
	if !ctx.IsSet(client.KeyAddress) && len(config.API.Address) > 0 {
		// set the API address field to value from config
		err = ctx.Set(client.KeyAddress, config.API.Address)
		if err != nil {
			return err
		}
	}

	// check if the API token is set in the context
	if !ctx.IsSet(client.KeyToken) && len(config.API.Token) > 0 {
		// set the API token field to value from config
		err = ctx.Set(client.KeyToken, config.API.Token)
		if err != nil {
			return err
		}
	}

	// check if the API version is set in the context
	if !ctx.IsSet("api.version") && len(config.API.Version) > 0 {
		// set the API version field to value from config
		err = ctx.Set("api.version", config.API.Version)
		if err != nil {
			return err
		}
	}

	// check if the log level is set in the context
	if !ctx.IsSet("log.level") && len(config.Log.Level) > 0 {
		// set the log level field to value from config
		err = ctx.Set("log.level", config.Log.Level)
		if err != nil {
			return err
		}
	}

	// check if the output is set in the context
	if !ctx.IsSet("output") && len(config.Output) > 0 {
		// set the output field to value from config
		err = ctx.Set("output", config.Output)
		if err != nil {
			return err
		}
	}

	// check if the org is set in the context
	if !ctx.IsSet("org") && len(config.Org) > 0 {
		// set the org field to value from config
		err = ctx.Set("org", config.Org)
		if err != nil {
			return err
		}
	}

	// check if the repo is set in the context
	if !ctx.IsSet("repo") && len(config.Repo) > 0 {
		// set the repo field to value from config
		err = ctx.Set("repo", config.Repo)
		if err != nil {
			return err
		}
	}

	// check if the secret engine is set in the context
	if !ctx.IsSet("secret.engine") && len(config.Secret.Engine) > 0 {
		// set the secret engine field to value from config
		err = ctx.Set("secret.engine", config.Secret.Engine)
		if err != nil {
			return err
		}
	}

	// check if the secret type is set in the context
	if !ctx.IsSet("secret.type") && len(config.Secret.Type) > 0 {
		// set the secret type field to value from config
		err = ctx.Set("secret.type", config.Secret.Type)
		if err != nil {
			return err
		}
	}

	return nil
}
