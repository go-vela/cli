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

func TestSchedule_Config_Update(t *testing.T) {
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
			name:    "success with dump output",
			failure: false,
			config: &Config{
				Action: "update",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "dump",
				Branch: "main",
			},
		},
		{
			name:    "success with json output",
			failure: false,
			config: &Config{
				Action: "update",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "json",
				Branch: "main",
			},
		},
		{
			name:    "success with spew output",
			failure: false,
			config: &Config{
				Action: "update",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "spew",
				Branch: "main",
			},
		},
		{
			name:    "success with yaml output",
			failure: false,
			config: &Config{
				Action: "update",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "yaml",
				Branch: "main",
			},
		},
		{
			name:    "failure",
			failure: true,
			config: &Config{
				Action: "update",
				Org:    "github",
				Repo:   "octocat",
				Name:   "not-found",
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err = test.config.Update(client)

			if test.failure {
				if err == nil {
					t.Errorf("Update for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Update for %s returned err: %v", test.name, err)
			}
		})
	}
}
