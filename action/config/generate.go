// SPDX-License-Identifier: Apache-2.0

package config

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v3"
)

// create filesystem based on the operating system
//
// https://godoc.org/github.com/spf13/afero#NewOsFs
var appFS = afero.NewOsFs()

// Generate produces a config file based off the provided configuration.
func (c *Config) Generate() error {
	logrus.Debug("executing generate for config file configuration")

	// create the config file content
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#ConfigFile
	config := &ConfigFile{
		API: &API{
			Address:      c.Addr,
			Token:        c.Token,
			AccessToken:  c.AccessToken,
			RefreshToken: c.RefreshToken,
			Version:      c.Version,
		},
		Log: &Log{
			Level: c.LogLevel,
		},
		Secret: &Secret{
			Engine: c.Engine,
			Type:   c.Type,
		},
		Compiler: &Compiler{
			GitHub: &GitHub{
				Token: c.GitHub.Token,
				URL:   c.GitHub.URL,
			},
		},
		Org:         c.Org,
		Repo:        c.Repo,
		Output:      c.Output,
		Color:       &c.Color.Enabled,
		ColorFormat: c.Color.Format,
		ColorTheme:  c.Color.Theme,
		NoGit:       c.NoGit,
	}

	logrus.Trace("creating file content for config file")

	// create output for config file
	//
	// https://pkg.go.dev/gopkg.in/yaml.v3?tab=doc#Marshal
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

	logrus.Tracef("creating directory structure to %s", c.File)

	// send Filesystem call to create directory path for config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#OsFs.MkdirAll
	err = a.Fs.MkdirAll(filepath.Dir(c.File), 0777)
	if err != nil {
		return err
	}

	logrus.Tracef("writing file content to %s", c.File)

	// send Filesystem call to create config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.WriteFile
	return a.WriteFile(c.File, out, 0600)
}
