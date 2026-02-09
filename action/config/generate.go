// SPDX-License-Identifier: Apache-2.0

package config

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	yaml "go.yaml.in/yaml/v3"
)

// create filesystem based on the operating system
//
// https://godoc.org/github.com/spf13/afero#NewOsFs
var appFS = afero.NewOsFs()

// Generate produces a config file based off the provided configuration.
func (c *Config) Generate() error {
	logrus.Debug("executing generate for config file configuration")

	out, err := genBytes(c)
	if err != nil {
		return err
	}

	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	if c.UseMemMap {
		a.Fs = afero.NewMemMapFs()
	}

	logrus.Tracef("creating directory structure to %s", c.File)

	// send Filesystem call to create directory path for config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#OsFs.MkdirAll
	err = a.MkdirAll(filepath.Dir(c.File), 0777)
	if err != nil {
		return err
	}

	logrus.Tracef("writing file content to %s", c.File)

	// send Filesystem call to create config file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.WriteFile
	return a.WriteFile(c.File, out, 0600)
}

func genBytes(c *Config) ([]byte, error) {
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
		Org:    c.Org,
		Repo:   c.Repo,
		Output: c.Output,
		Color:  &c.Color.Enabled,
		NoGit:  c.NoGit,
	}

	// only save if theme was user specified; this prevents saving default "monokai"
	// during login, which would bypass auto-detection for light terminal backgrounds.
	if c.Color.UserSpecified {
		config.ColorTheme = c.Color.Theme
	}

	// only save if not default
	if c.Color.Format != "" && c.Color.Format != "terminal256" {
		config.ColorFormat = c.Color.Format
	}

	out, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	return out, nil
}
