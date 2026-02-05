// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestDashboard_Config_View(t *testing.T) {
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
				ID:     "c8da1302-07d6-11ea-882f-4893bca275b8",
				Full:   true,
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				ID:     "c8da1302-07d6-11ea-882f-4893bca275b8",
				Full:   false,
				Output: "dump",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.View(t.Context(), client)

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
