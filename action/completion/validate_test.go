// SPDX-License-Identifier: Apache-2.0

package completion

import (
	"testing"
)

func TestCompletion_Config_Validate(t *testing.T) {
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
				Bash:   true,
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
		err := test.config.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}
