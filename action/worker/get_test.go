// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestWorker_Config_Get(t *testing.T) {
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
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Output: "wide",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "get",
				Output: "yaml",
			},
		},
		// TODO: mock doesn't have failure for listing workers
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
