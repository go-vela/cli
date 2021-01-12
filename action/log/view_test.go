// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestLog_Config_ViewService(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
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
				Action:  "view",
				Org:     "github",
				Repo:    "octocat",
				Build:   1,
				Service: 1,
				Output:  "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "view",
				Org:     "github",
				Repo:    "octocat",
				Build:   1,
				Service: 1,
				Output:  "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "view",
				Org:     "github",
				Repo:    "octocat",
				Build:   1,
				Service: 1,
				Output:  "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "view",
				Org:     "github",
				Repo:    "octocat",
				Build:   1,
				Service: 1,
				Output:  "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "view",
				Org:     "github",
				Repo:    "octocat",
				Build:   1,
				Service: 1,
				Output:  "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.ViewService(client)

		if test.failure {
			if err == nil {
				t.Errorf("ViewService should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("ViewService returned err: %v", err)
		}
	}
}

func TestService_Config_ViewStep(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
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
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Step:   1,
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Step:   1,
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Step:   1,
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Step:   1,
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Step:   1,
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.ViewStep(client)

		if test.failure {
			if err == nil {
				t.Errorf("ViewStep should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("ViewStep returned err: %v", err)
		}
	}
}
