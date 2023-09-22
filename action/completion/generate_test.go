// SPDX-License-Identifier: Apache-2.0

package completion

import (
	"testing"
)

func TestCompletion_Config_Generate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "generate",
				Bash:   true,
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "generate",
				Zsh:    true,
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "generate",
			},
		},
	}

	// run tests
	for _, test := range tests {
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
