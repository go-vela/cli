// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestLog_Config_Get(t *testing.T) {
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
				Action: "get",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Org:    "github",
				Repo:   "octocat",
				Build:  1,
				Output: "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Get(client)

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
