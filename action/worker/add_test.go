// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestWorker_Config_Add(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	// set up new gin instance for fake worker
	e := gin.New()

	// mock endpoint for worker register call
	e.POST("/register", func(c *gin.Context) { c.JSON(http.StatusOK, "worker registered successfully") })

	// create a new test server
	w := httptest.NewServer(e)

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
		// TODO: mock doesn't have failure for worker creation
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
