// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_Config_Generate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
			},
		},
	}

	// run tests
	for _, test := range tests {
		// setup filesystem
		appFS = afero.NewMemMapFs()

		err := test.config.Generate()

		if test.failure {
			if err == nil {
				t.Errorf("Generate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Generate returned err: %v", err)
		}
	}
}
