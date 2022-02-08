// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"testing"

	"github.com/go-vela/cli/internal"
)

func TestRepo_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:     internal.ActionAdd,
				Org:        "github",
				Name:       "octocat",
				Link:       "https://github.com/github/octocat",
				Clone:      "https://github.com/github/octocat.git",
				Branch:     "master",
				Timeout:    60,
				Visibility: "public",
				Private:    false,
				Trusted:    false,
				Active:     true,
				Events:     []string{"push", "pull_request", "comment", "deployment", "tag"},
				Output:     "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionChown,
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  internal.ActionGet,
				Org:     "github",
				Page:    1,
				PerPage: 10,
				Output:  "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionRemove,
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionRepair,
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:     internal.ActionUpdate,
				Org:        "github",
				Name:       "octocat",
				Link:       "https://github.com/github/octocat",
				Clone:      "https://github.com/github/octocat.git",
				Branch:     "master",
				Timeout:    60,
				Visibility: "public",
				Private:    false,
				Trusted:    false,
				Active:     true,
				Events:     []string{"push", "pull_request", "comment", "deployment", "tag"},
				Output:     "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionView,
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
				Org:    "",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
				Org:    "github",
				Name:   "",
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
