// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
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
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "remove",
				File:   "testdata/config.yml",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				File:   "testdata/config.yml",
				RemoveFlags: []string{
					"api.addr",
					"token",
					"api.version",
					"log.level",
					"engine",
					"type",
					"org",
					"repo",
					"output",
				},
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
