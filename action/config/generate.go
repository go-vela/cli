// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"path/filepath"

	"github.com/spf13/afero"

	yaml "gopkg.in/yaml.v2"
)

// create filesystem based on the operating system
var appFS = afero.NewOsFs()

// Generate produces a pipeline based off the provided configuration.
func (c *Config) Generate() error {
	// create the config file content
	config := &ConfigFile{
		API: &API{
			Address: c.Addr,
			Token:   c.Token,
			Version: c.Version,
		},
		Log: &Log{
			Level: c.LogLevel,
		},
		Secret: &Secret{
			Engine: c.Engine,
			Type:   c.Type,
		},
		Org:  c.Org,
		Repo: c.Repo,
		Output: c.Output,
	}

	// create output for config file
	out, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// send Filesystem call to create directory path for config file
	err = a.Fs.MkdirAll(filepath.Dir(c.File), 0777)
	if err != nil {
		return err
	}

	// send Filesystem call to create config file
	return a.WriteFile(c.File, []byte(out), 0600)
}
