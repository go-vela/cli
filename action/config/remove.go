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

// Remove deletes one or more fields from the config file based off the provided configuration.
func (c *Config) Remove() error {
	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if remove flags are empty
	if len(c.RemoveFlags) == 0 {
		// send Filesystem call to delete config file
		return a.Remove(c.File)
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

	// iterate through all flags to be removed
	for _, flag := range c.RemoveFlags {
		// check if API addr flag should be removed
		if strings.EqualFold(flag, client.KeyAddress) {
			// set the API addr field to empty in config
			config.API.Address = ""
		}

		// check if API token flag should be removed
		if strings.EqualFold(flag, client.KeyToken) {
			// set the API token field to empty in config
			config.API.Token = ""
		}

		// check if API version flag should be removed
		if strings.EqualFold(flag, "version") {
			// set the API version field to empty in config
			config.API.Version = ""
		}

		// check if log level flag should be removed
		if strings.EqualFold(flag, "level") {
			// set the log level field to empty in config
			config.Log.Level = ""
		}

		// check if secret engine flag should be removed
		if strings.EqualFold(flag, "engine") {
			// set the secret engine field to empty in config
			config.Secret.Engine = ""
		}

		// check if secret type flag should be removed
		if strings.EqualFold(flag, "type") {
			// set the secret type field to empty in config
			config.Secret.Type = ""
		}

		// check if org flag should be removed
		if strings.EqualFold(flag, "org") {
			// set the org field to empty in config
			config.Org = ""
		}

		// check if repo flag should be removed
		if strings.EqualFold(flag, "repo") {
			// set the repo field to empty in config
			config.Repo = ""
		}

		// check if output flag should be removed
		if strings.EqualFold(flag, "output") {
			// set the output field to empty in config
			config.Output = ""
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
