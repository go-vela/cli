// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package schedule

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSchedule_Config_Get(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create test client: %v", err)
	}

	// setup tests
	tests := []struct {
		name    string
		failure bool
		config  *Config
	}{
		{
			name:    "success with empty output",
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
			name:    "success with dump output",
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "dump",
			},
		},
		{
			name:    "success with json output",
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "json",
			},
		},
		{
			name:    "success with spew output",
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "spew",
			},
		},
		{
			name:    "success with wide output",
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "wide",
			},
		},
		{
			name:    "success with yaml output",
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err = test.config.Get(client)

			if test.failure {
				if err == nil {
					t.Errorf("Get for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Get for %s returned err: %v", test.name, err)
			}
		})
	}
}
