// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestRepo_Config_Remove(t *testing.T) {
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
				Org:    "github",
				Name:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Name:   "octocat",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Name:   "octocat",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Name:   "octocat",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Name:   "octocat",
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "remove",
				Org:    "github",
				Name:   "not-found",
				Output: "",
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
