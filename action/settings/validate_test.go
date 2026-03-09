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
				Action: "update",
				Compiler: Compiler{
					CloneImage: new("test"),
				},
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Compiler: Compiler{
					TemplateDepth: new(1),
				},
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Compiler: Compiler{
					StarlarkExecLimit: new(int64(1)),
				},
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Queue: Queue{
					Routes: &[]string{"test"},
				},
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Queue: Queue{
					Routes: new([]string{"test"}),
				},
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
