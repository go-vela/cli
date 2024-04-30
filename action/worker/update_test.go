// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestWorker_Config_Update(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	tBool := true

	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:     "update",
				Hostname:   "MyWorker",
				Address:    "myworker.example.com",
				Routes:     []string{"large", "small"},
				BuildLimit: 1,
				Active:     &tBool,
				Output:     "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:     "update",
				Hostname:   "MyWorker",
				Address:    "myworker.example.com",
				Routes:     []string{"large", "small"},
				BuildLimit: 1,
				Active:     &tBool,
				Output:     "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:     "update",
				Hostname:   "MyWorker",
				Address:    "myworker.example.com",
				Routes:     []string{"large", "small"},
				BuildLimit: 1,
				Active:     &tBool,
				Output:     "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:     "update",
				Hostname:   "MyWorker",
				Address:    "myworker.example.com",
				Routes:     []string{"large", "small"},
				BuildLimit: 1,
				Active:     &tBool,
				Output:     "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:     "update",
				Hostname:   "MyWorker",
				Address:    "myworker.example.com",
				Routes:     []string{"large", "small"},
				BuildLimit: 1,
				Active:     &tBool,
				Output:     "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:     "update",
				Hostname:   "0",
				Address:    "myworker.example.com",
				Routes:     []string{"large", "small"},
				BuildLimit: 1,
				Active:     &tBool,
				Output:     "",
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
