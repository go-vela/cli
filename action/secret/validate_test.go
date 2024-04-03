// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"testing"
)

func TestSecret_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "add",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				File:   "testdata/repo.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "get",
				Engine:  "native",
				Type:    "repo",
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
				Action: "remove",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Engine:      "native",
				Type:        "repo",
				Org:         "github",
				Repo:        "octocat",
				Name:        "foo",
				Value:       "bar",
				AllowEvents: []string{"comment", "push", "pull_request", "tag", "deployment", "schedule"},
				Output:      "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:      "add",
				Engine:      "native",
				Type:        "repo",
				Org:         "github",
				Repo:        "octocat",
				Name:        "foo",
				Value:       "bar",
				AllowEvents: []string{"foo"},
				Output:      "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "add",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Engine: "",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Engine: "native",
				Type:   "",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Engine: "native",
				Type:   "baz",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Engine: "native",
				Type:   "repo",
				Org:    "",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Engine: "native",
				Type:   "shared",
				Org:    "github",
				Team:   "",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
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
