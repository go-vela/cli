// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestWorker_Config_Add(t *testing.T) {
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
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "myworker.example.com",
				Output:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "myworker.example.com",
				Output:   "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "myworker.example.com",
				Output:   "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "myworker.example.com",
				Output:   "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "myworker.example.com",
				Output:   "yaml",
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
