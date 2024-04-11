// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSecret_Config_Update(t *testing.T) {
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
				Engine: "native",
				Type:   "org",
				Org:    "github",
				Repo:   "*",
				Name:   "foo",
				Value:  "bar",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "shared",
				Org:    "github",
				Team:   "octokitties",
				Name:   "foo",
				Value:  "bar",
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
				Value:  "@testdata/foo.txt",
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
				Output: "dump",
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
				Output: "json",
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
				Output: "spew",
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
				Output: "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Update(client)

		if test.failure {
			if err == nil {
				t.Errorf("Update should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Update returned err: %v", err)
		}
	}
}

func TestSecret_Config_UpdateFromFile(t *testing.T) {
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
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/org.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/shared.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/multiple.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.UpdateFromFile(client)

		if test.failure {
			if err == nil {
				t.Errorf("UpdateFromFile should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("UpdateFromFile returned err: %v", err)
		}
	}
}
