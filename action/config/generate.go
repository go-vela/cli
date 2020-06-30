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
//
// https://godoc.org/github.com/spf13/afero#NewOsFs
var appFS = afero.NewOsFs()

// Generate produces a pipeline based off the provided configuration.
func (c *Config) Generate() error {
	// create the config file content
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#ConfigFile
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
		Org:    c.Org,
		Repo:   c.Repo,
		Output: c.Output,
	}

	// create output for config file
	//
	// https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc#Marshal
	out, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	// send Filesystem call to create directory path for config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#OsFs.MkdirAll
	err = a.Fs.MkdirAll(filepath.Dir(c.File), 0777)
	if err != nil {
		return err
	}

	// send Filesystem call to create config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.WriteFile
	return a.WriteFile(c.File, []byte(out), 0600)
}
