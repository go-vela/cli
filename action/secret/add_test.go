// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestSecret_Config_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	// setup tests
	tests := []struct {
		config  *Config
		failure bool
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
				Engine: "native",
				Type:   "org",
				Org:    "github",
				Name:   "foo",
				Value:  "bar",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
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
				Action: "add",
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
				Action: "add",
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
				Action: "add",
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
				Action: "add",
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
				Action: "add",
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

func TestSecret_Config_AddFromFile(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	// setup tests
	tests := []struct {
		config  *Config
		failure bool
	}{
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
				Action: "add",
				File:   "testdata/org.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				File:   "testdata/shared.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				File:   "testdata/multiple.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				File:   "testdata/repo.yml",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				File:   "testdata/repo.yml",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				File:   "testdata/repo.yml",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				File:   "testdata/repo.yml",
				Output: "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.AddFromFile(client)

		if test.failure {
			if err == nil {
				t.Errorf("AddFromFile should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("AddFromFile returned err: %v", err)
		}
	}
}
