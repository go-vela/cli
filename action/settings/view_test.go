// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestWorker_Config_View(t *testing.T) {
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
				Action:   "view",
				Hostname: "MyWorker",
				Output:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "view",
				Hostname: "MyWorker",
				Output:   "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "view",
				Hostname: "MyWorker",
				Output:   "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "view",
				Hostname: "MyWorker",
				Output:   "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "view",
				Hostname: "MyWorker",
				Output:   "yaml",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:            "view",
				Hostname:          "MyWorker",
				RegistrationToken: true,
				Output:            "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "view",
				Hostname: "0",
				Output:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:            "view",
				Hostname:          "",
				Output:            "",
				RegistrationToken: true,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.View(client)

		if test.failure {
			if err == nil {
				t.Errorf("View should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("View returned err: %v", err)
		}
	}
}
