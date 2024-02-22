// SPDX-License-Identifier: Apache-2.0

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
				Branch:     "main",
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
			failure: true,
			config: &Config{
				Action:       "add",
				Org:          "github",
				Name:         "octocat",
				ApproveBuild: "invalid",
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "chown",
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Page:    1,
				PerPage: 10,
				Output:  "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "repair",
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:       "update",
				Org:          "github",
				Name:         "octocat",
				Link:         "https://github.com/github/octocat",
				Clone:        "https://github.com/github/octocat.git",
				Branch:       "main",
				Timeout:      60,
				Visibility:   "public",
				Private:      false,
				Trusted:      false,
				Active:       true,
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
				ApproveBuild: "fork-no-write",
				Output:       "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:       "update",
				Org:          "github",
				Name:         "octocat",
				Link:         "https://github.com/github/octocat",
				Clone:        "https://github.com/github/octocat.git",
				Branch:       "main",
				Timeout:      60,
				Visibility:   "public",
				Private:      false,
				Trusted:      false,
				Active:       true,
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
				ApproveBuild: "first-time",
				Output:       "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:       "update",
				Org:          "github",
				Name:         "octocat",
				ApproveBuild: "invalid",
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
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
