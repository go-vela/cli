// SPDX-License-Identifier: Apache-2.0

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"

	"github.com/go-vela/cli/internal/output"
)

// View inspects the config file based off the provided configuration.
func (c *Config) View() error {
	logrus.Debug("executing view for config file configuration")

	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	if c.UseMemMap {
		a.Fs = afero.NewMemMapFs()

		bytes, err := a.ReadFile(c.File)
		if err != nil || len(bytes) == 0 {
			err = writeTestConfig(a, c)
			if err != nil {
				return err
			}
		}
	}

	logrus.Tracef("reading content from %s", c.File)

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
	return output.Stdout(string(config))
}
