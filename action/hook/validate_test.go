// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

import (
	"testing"
)

func TestHook_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		config  *Config
		failure bool
	}{
		{
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "",
				Repo:   "octocat",
				Number: 1,
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "",
				Number: 1,
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
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
