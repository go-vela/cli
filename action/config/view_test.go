// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_Config_View(t *testing.T) {
	// setup tests
	tests := []struct {
		config  *Config
		failure bool
	}{
		{
			failure: false,
			config: &Config{
				Action: "view",
				File:   "testdata/config.yml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				File:   "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		// setup filesystem
		appFS = afero.NewOsFs()

		err := test.config.View()

		if test.failure {
			if err == nil {
				t.Errorf("View should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("View returned err: %v", err)
		}
	}
}
