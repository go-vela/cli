// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestWorker_Config_Get(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	tBool := true
	fBool := false

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
				Action:          "get",
				Output:          "",
				CheckedInBefore: 1,
				CheckedInAfter:  0,
				Active:          &tBool,
			},
		},
		{
			failure: false,
			config: &Config{
				Action:          "get",
				Output:          "dump",
				CheckedInBefore: 1,
				Active:          &fBool,
			},
		},
		{
			failure: false,
			config: &Config{
				Action:         "get",
				Output:         "json",
				CheckedInAfter: 0,
				Active:         &tBool,
			},
		},
		{
			failure: false,
			config: &Config{
				Action:          "get",
				Output:          "spew",
				CheckedInBefore: 1,
				CheckedInAfter:  0,
				Active:          &fBool,
			},
		},
		{
			failure: false,
			config: &Config{
				Action:          "get",
				Output:          "wide",
				CheckedInBefore: 1,
				CheckedInAfter:  0,
			},
		},
		{
			failure: false,
			config: &Config{
				Action:          "get",
				Output:          "yaml",
				CheckedInBefore: 1,
				CheckedInAfter:  0,
				Active:          &tBool,
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
