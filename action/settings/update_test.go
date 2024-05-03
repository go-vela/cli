// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSettings_Config_Update(t *testing.T) {
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
				Action: "update",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "update",
				Output: "",
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
