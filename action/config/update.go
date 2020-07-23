// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"strings"

	"github.com/go-vela/cli/internal"

	"github.com/spf13/afero"

	yaml "gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"
)

// Update modifies one or more fields from the config file based off the provided configuration.
func (c *Config) Update() error {
	logrus.Debug("executing update for config file configuration")

	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	logrus.Tracef("reading content from %s", c.File)

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

	// iterate through all flags to be modified
	for key, value := range c.UpdateFlags {
		logrus.Tracef("updating key %s with value %s", key, value)

		// check if API address flag should be modified
		if strings.EqualFold(key, internal.FlagAPIAddress) {
			// set the API address field to value provided
			config.API.Address = value
		}

		// check if API token flag should be modified
		if strings.EqualFold(key, internal.FlagAPIToken) {
			// set the API token field to value provided
			config.API.Token = value
		}

		// check if API version flag should be modified
		if strings.EqualFold(key, internal.FlagAPIVersion) {
			// set the API version field to value provided
			config.API.Version = value
		}

		// check if log level flag should be modified
		if strings.EqualFold(key, internal.FlagLogLevel) {
			// set the log level field to value provided
			config.Log.Level = value
		}

		// check if secret engine flag should be modified
		if strings.EqualFold(key, internal.FlagSecretEngine) {
			// set the secret engine field to value provided
			config.Secret.Engine = value
		}

		// check if secret type flag should be modified
		if strings.EqualFold(key, internal.FlagSecretType) {
			// set the secret type field to value provided
			config.Secret.Type = value
		}

		// check if org flag should be modified
		if strings.EqualFold(key, internal.FlagOrg) {
			// set the org field to value provided
			config.Org = value
		}

		// check if repo flag should be modified
		if strings.EqualFold(key, internal.FlagRepo) {
			// set the repo field to value provided
			config.Repo = value
		}

		// check if output flag should be modified
		if strings.EqualFold(key, internal.FlagOutput) {
			// set the output field to value provided
			config.Output = value
		}
	}

	logrus.Trace("creating file content for config file")

	// create output for config file
	//
	// https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc#Marshal
	out, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	logrus.Tracef("writing file content to %s", c.File)

	// send Filesystem call to create config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.WriteFile
	return a.WriteFile(c.File, []byte(out), 0600)
}
