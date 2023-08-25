// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package schedule

import (
	"testing"
)

func TestSchedule_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		name    string
		failure bool
		config  *Config
	}{
		{
			name:    "success with add",
			failure: false,
			config: &Config{
				Action: "add",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "",
				Branch: "main",
			},
		},
		{
			name:    "success with get",
			failure: false,
			config: &Config{
				Action: "get",
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			name:    "success with remove",
			failure: false,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			name:    "success with update",
			failure: false,
			config: &Config{
				Action: "update",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "",
				Branch: "main",
			},
		},
		{
			name:    "success with view",
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			name:    "failure with no org",
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
		{
			name:    "failure with no repo",
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "",
				Name:   "foo",
				Output: "",
			},
		},
		{
			name:    "failure with no name",
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Name:   "",
				Output: "",
			},
		},
		{
			name:    "failure with no entry",
			failure: true,
			config: &Config{
				Action: "add",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.config.Validate()

			if test.failure {
				if err == nil {
					t.Errorf("Validate for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Validate for %s returned err: %v", test.name, err)
			}
		})
	}
}
