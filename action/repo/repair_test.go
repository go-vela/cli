// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestRepo_Config_Repair(t *testing.T) {
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
				Action: "repair",
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "repair",
				Org:    "github",
				Name:   "octocat",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "repair",
				Org:    "github",
				Name:   "octocat",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "repair",
				Org:    "github",
				Name:   "octocat",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "repair",
				Org:    "github",
				Name:   "octocat",
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "repair",
				Org:    "github",
				Name:   "not-found",
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Repair(client)

		if test.failure {
			if err == nil {
				t.Errorf("Repair should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Repair returned err: %v", err)
		}
	}
}
