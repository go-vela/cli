// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"strings"

	"github.com/go-vela/cli/internal/client"

	"github.com/spf13/afero"

	yaml "gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"
)

// Remove deletes one or more fields from the config file based off the provided configuration.
func (c *Config) Remove() error {
	logrus.Debug("executing remove for config file configuration")

	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if remove flags are empty
	if len(c.RemoveFlags) == 0 {
		logrus.Tracef("removing config file %s", c.File)

		// send Filesystem call to delete config file
		//
		// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.Remove
		return a.Remove(c.File)
	}

	logrus.Tracef("reading content from %s", c.File)

	// send Filesystem call to read config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.ReadFile
	data, err := a.ReadFile(c.File)
	if err != nil {
		return err
	}

	// create the config object
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

	// iterate through all flags to be removed
	for _, flag := range c.RemoveFlags {
		logrus.Tracef("removing key %s", flag)

		// check if API address flag should be removed
		if strings.EqualFold(flag, client.KeyAddress) {
			// set the API address field to empty in config
			config.API.Address = ""
		}

		// check if API token flag should be removed
		if strings.EqualFold(flag, client.KeyToken) {
			// set the API token field to empty in config
			config.API.Token = ""
		}

		// check if API version flag should be removed
		if strings.EqualFold(flag, "api.version") {
			// set the API version field to empty in config
			config.API.Version = ""
		}

		// check if log level flag should be removed
		if strings.EqualFold(flag, "log.level") {
			// set the log level field to empty in config
			config.Log.Level = ""
		}

		// check if secret engine flag should be removed
		if strings.EqualFold(flag, "secret.engine") {
			// set the secret engine field to empty in config
			config.Secret.Engine = ""
		}

		// check if secret type flag should be removed
		if strings.EqualFold(flag, "secret.type") {
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
