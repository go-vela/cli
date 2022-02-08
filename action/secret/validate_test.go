// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"testing"

	"github.com/go-vela/cli/internal"
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
				Action: internal.ActionAdd,
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
				Action: internal.ActionAdd,
				File:   "testdata/repo.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  internal.ActionGet,
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
				Action: internal.ActionRemove,
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
				Action: internal.ActionUpdate,
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
				Action: internal.ActionUpdate,
				File:   "testdata/repo.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionView,
				Engine: "native",
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
				Action: internal.ActionAdd,
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Events: []string{"foo"},
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionAdd,
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
				Action: internal.ActionView,
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
				Action: internal.ActionView,
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
				Action: internal.ActionView,
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
				Action: internal.ActionView,
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
				Action: internal.ActionView,
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
				Action: internal.ActionView,
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
				Action: internal.ActionView,
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
