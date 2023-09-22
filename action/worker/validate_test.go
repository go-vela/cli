// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"testing"
)

func TestWorker_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "myworker.example.com",
				Output:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "",
				Output:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "",
				Address:  "myworker.example.com",
				Output:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "",
				Output:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "::",
				Output:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "view",
				Hostname: "MyWorker",
				Output:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "view",
				Hostname: "",
				Output:   "",
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
				Action:     "update",
				Hostname:   "MyWorker",
				Address:    "myworker.example.com",
				Active:     true,
				BuildLimit: 1,
				Routes:     []string{"large", "small"},
				Output:     "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:     "update",
				Hostname:   "",
				Address:    "myworker.example.com",
				Active:     true,
				BuildLimit: 1,
				Routes:     []string{"large", "small"},
				Output:     "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:     "update",
				Hostname:   "",
				Address:    "::",
				Active:     true,
				BuildLimit: 1,
				Routes:     []string{"large", "small"},
				Output:     "",
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
