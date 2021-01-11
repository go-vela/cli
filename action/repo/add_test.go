// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestRepo_Config_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "Vela CLI", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

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
				Output:     "",
			},
		},
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
				Output:     "dump",
			},
		},
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
				Output:     "json",
			},
		},
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
				Output:     "spew",
			},
		},
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
				Output:     "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Add(client)

		if test.failure {
			if err == nil {
				t.Errorf("Add should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Add returned err: %v", err)
		}
	}
}
