// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestRepo_Config_Sync(t *testing.T) {
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
				Action: "sync",
				Org:    "github",
				Name:   "octocat",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "sync",
				Org:    "github",
				Name:   "not-found",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Sync(client)

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

func TestRepo_Config_SyncAll(t *testing.T) {
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
				Action: "sync",
				Org:    "github",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "sync",
				Org:    "not-found",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.SyncAll(client)

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
