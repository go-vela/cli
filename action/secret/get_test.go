// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSecret_Config_Get(t *testing.T) {
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
				Action:  "get",
				Engine:  "native",
				Type:    "org",
				Org:     "github",
				Repo:    "*",
				Page:    1,
				PerPage: 10,
				Output:  "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:  "get",
				Engine:  "native",
				Type:    "shared",
				Org:     "github",
				Team:    "octokitties",
				Name:    "foo",
				Page:    1,
				PerPage: 10,
				Output:  "",
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
				Output:  "dump",
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
				Output:  "json",
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
				Output:  "spew",
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
				Output:  "wide",
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
				Output:  "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Get(t.Context(), client)

		if test.failure {
			if err == nil {
				t.Errorf("Get should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Get returned err: %v", err)
		}
	}
}
