// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_Config_Update(t *testing.T) {
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
				GitHub: &GitHub{},
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				File:   "testdata/config.yml",
				UpdateFlags: map[string]string{
					"api.addr":             "https://vela-server.localhost",
					"api.token":            "superSecretToken",
					"api.token.access":     "superSecretAccessToken",
					"api.token.refresh":    "superSecretRefreshToken",
					"api.version":          "1",
					"log.level":            "info",
					"no-git":               "true",
					"secret.engine":        "native",
					"secret.type":          "repo",
					"compiler.GitHubToken": "somePATToken",
					"compiler.GitHubURL":   "github.com",
					"org":                  "github",
					"repo":                 "octocat",
					"output":               "json",
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
			GitHub: &GitHub{},
		}

		// generate config file
		err := config.Generate()
		if err != nil {
			t.Errorf("unable to generate config: %v", err)
		}

		err = test.config.Update()

		if test.failure {
			if err == nil {
				t.Errorf("Update should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Update returned err: %v", err)
		}
	}
}
