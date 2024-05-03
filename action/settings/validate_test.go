// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"testing"
)

func TestSettings_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "add",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "add",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "add",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "add",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "update",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "update",
				Output: "",
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
