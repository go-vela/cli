// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_Config_Remove(t *testing.T) {
	// setup tests
	tests := []struct {
		config  *Config
		failure bool
	}{
		{
			failure: false,
			config: &Config{
				Action: "remove",
				File:   "testdata/config.yml",
				GitHub: &GitHub{},
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				File:   "testdata/config.yml",
				RemoveFlags: []string{
					"api.addr",
					"api.token",
					"api.token.access",
					"api.token.refresh",
					"api.version",
					"log.level",
					"no-git",
					"secret.engine",
					"secret.type",
					"compiler.GitHubToken",
					"compiler.GitHubURL",
					"org",
					"repo",
					"output",
				},
				GitHub: &GitHub{},
			},
		},
	}

	// run tests
	for _, test := range tests {
		// setup filesystem
		appFS = afero.NewMemMapFs()

		// create test config for generating file
		config := &Config{
			Action: "generate",
			File:   test.config.File,
			GitHub: &GitHub{},
		}

		// generate config file
		err := config.Generate()
		if err != nil {
			t.Errorf("unable to generate config: %v", err)
		}

		err = test.config.Remove()

		if test.failure {
			if err == nil {
				t.Errorf("Remove should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Remove returned err: %v", err)
		}
	}
}
