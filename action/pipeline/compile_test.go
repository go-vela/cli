// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestPipeline_Config_Compile(t *testing.T) {
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
				Action: "compile",
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "compile",
				Org:    "github",
				Repo:   "octocat",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "compile",
				Org:    "github",
				Repo:   "octocat",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "compile",
				Org:    "github",
				Repo:   "octocat",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "compile",
				Org:    "github",
				Repo:   "octocat",
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "compile",
				Org:    "github",
				Repo:   "not-found",
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Compile(client)

		if test.failure {
			if err == nil {
				t.Errorf("Compile should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Compile returned err: %v", err)
		}
	}
}
