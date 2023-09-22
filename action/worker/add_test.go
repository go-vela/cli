// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/worker/mock/worker"

	"github.com/go-vela/sdk-go/vela"
)

func TestWorker_Config_Add(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	// create mock server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	// create mock worker server
	w := httptest.NewServer(worker.FakeHandler())

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
				Address:  w.URL,
				Output:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  w.URL,
				Output:   "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  w.URL,
				Output:   "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  w.URL,
				Output:   "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  w.URL,
				Output:   "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "add",
				Hostname: "",
				Address:  w.URL,
				Output:   "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "add",
				Hostname: "MyWorker",
				Address:  "",
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
