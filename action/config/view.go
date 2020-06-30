// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/spf13/afero"
)

// View inspects the config file based off the provided configuration.
func (c *Config) View() error {
	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	// send Filesystem call to read config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.ReadFile
	config, err := a.ReadFile(c.File)
	if err != nil {
		return err
	}

	// output the config in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(config)
}
