// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package service

import (
	"testing"

	"github.com/go-vela/cli/internal"
)

func TestService_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:  internal.ActionGet,
				Org:     "github",
				Repo:    "octocat",
				Build:   1,
				Page:    1,
				PerPage: 10,
				Output:  "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionView,
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Number: 1,
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
				Org:    "",
				Repo:   "octocat",
				Build:  1,
				Number: 1,
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
				Org:    "github",
				Repo:   "",
				Build:  1,
				Number: 1,
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
				Org:    "github",
				Repo:   "octocat",
				Build:  0,
				Number: 1,
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Number: 0,
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
