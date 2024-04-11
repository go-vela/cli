// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSecret_Config_Remove(t *testing.T) {
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
				Action: "remove",
				Engine: "native",
				Type:   "org",
				Org:    "github",
				Repo:   "*",
				Name:   "foo",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				Engine: "native",
				Type:   "shared",
				Org:    "github",
				Team:   "octokitties",
				Name:   "foo",
				Output: "",
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
				Output: "dump",
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
				Output: "json",
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
				Output: "spew",
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
				Output: "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Remove(client)

		if test.failure {
			if err == nil {
				t.Errorf("Remove should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Remove returned err: %v", err)
		}
	}
}
