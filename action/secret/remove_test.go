// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestSecret_Config_Remove(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, nil)
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
				Output: "default",
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
				Output: "default",
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
				Output: "default",
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
