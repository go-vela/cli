// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSettings_Config_View(t *testing.T) {
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
				Action: "view",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Output: "yaml",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Output: "yaml",
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
