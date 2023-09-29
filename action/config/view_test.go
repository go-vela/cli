// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_Config_View(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
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
