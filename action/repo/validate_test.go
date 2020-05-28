// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"testing"
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
				Action:     "add",
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
				Output:     "default",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "chown",
				Org:    "github",
				Name:   "octocat",
				Output: "default",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Page:    1,
				PerPage: 10,
				Output:  "default",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Name:   "octocat",
				Output: "default",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "repair",
				Org:    "github",
				Name:   "octocat",
				Output: "default",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:     "update",
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
				Output:     "default",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Name:   "octocat",
				Output: "default",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "",
				Name:   "octocat",
				Output: "default",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Name:   "",
				Output: "default",
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
