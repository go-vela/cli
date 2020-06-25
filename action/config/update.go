// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"strings"

	"github.com/go-vela/cli/internal/client"

	"github.com/spf13/afero"

	yaml "gopkg.in/yaml.v2"
)

// Update modifies one or more fields from the config file based off the provided configuration.
func (c *Config) Update() error {
	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// send Filesystem call to read config file
	data, err := a.ReadFile(c.File)
	if err != nil {
		return err
	}

	// create the config object
	config := new(ConfigFile)

	// update the config object with the current content
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}

	// iterate through all flags to be modified
	for key, value := range c.UpdateFlags {
		// check if API addr flag should be modified
		if strings.EqualFold(key, client.KeyAddress) {
			// set the API addr field to value provided
			config.API.Address = value
		}

		// check if API token flag should be modified
		if strings.EqualFold(key, client.KeyToken) {
			// set the API token field to value provided
			config.API.Token = value
		}

		// check if API version flag should be modified
		if strings.EqualFold(key, "version") {
			// set the API version field to value provided
			config.API.Version = value
		}

		// check if log level flag should be modified
		if strings.EqualFold(key, "level") {
			// set the log level field to value provided
			config.Log.Level = value
		}

		// check if secret engine flag should be modified
		if strings.EqualFold(key, "engine") {
			// set the secret engine field to value provided
			config.Secret.Engine = value
		}

		// check if secret type flag should be modified
		if strings.EqualFold(key, "type") {
			// set the secret type field to value provided
			config.Secret.Type = value
		}

		// check if org flag should be modified
		if strings.EqualFold(key, "org") {
			// set the org field to value provided
			config.Org = value
		}

		// check if repo flag should be modified
		if strings.EqualFold(key, "repo") {
			// set the repo field to value provided
			config.Repo = value
		}

		// check if output flag should be modified
		if strings.EqualFold(key, "output") {
			// set the output field to value provided
			config.Output = value
		}
	}

	// create output for config file
	out, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// send Filesystem call to create config file
	return a.WriteFile(c.File, []byte(out), 0600)
}
