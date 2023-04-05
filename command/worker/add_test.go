// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"

	"github.com/urfave/cli/v2"
)

func TestWorker_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// set up new gin instance for fake worker
	e := gin.New()

	// mock endpoint for worker register call
	e.GET("/register", func(c *gin.Context) { c.JSON(http.StatusOK, "worker registered successfully") })

	// create a new test server
	w := httptest.NewServer(e)

	// setup flags
	authSet := flag.NewFlagSet("test", 0)
	authSet.String("api.addr", s.URL, "doc")
	authSet.String("api.token.access", test.TestTokenGood, "doc")
	authSet.String("api.token.refresh", "superSecretRefreshToken", "doc")

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token.access", test.TestTokenGood, "doc")
	fullSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	fullSet.String("worker.hostname", "MyWorker", "doc")
	fullSet.String("worker.address", w.URL, "doc")
	fullSet.String("output", "json", "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     fullSet,
		},
		{
			failure: true,
			set:     authSet,
		},
		{
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := add(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("add should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("add returned err: %v", err)
		}
	}
}
