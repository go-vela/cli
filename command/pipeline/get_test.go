// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"

	"github.com/urfave/cli/v2"
)

func TestPipeline_Get(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	authSet := flag.NewFlagSet("test", 0)
	authSet.String("api.addr", s.URL, "doc")
	authSet.String("api.token.access", test.TestTokenGood, "doc")
	authSet.String("api.token.refresh", "superSecretRefreshToken", "doc")

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token.access", test.TestTokenGood, "doc")
	fullSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	fullSet.String("org", "github", "doc")
	fullSet.String("repo", "octocat", "doc")
	fullSet.Int("page", 1, "doc")
	fullSet.Int("per.page", 10, "doc")
	fullSet.String("output", "json", "doc")

	// setup tests
	tests := []struct {
		name    string
		failure bool
		set     *flag.FlagSet
	}{
		{
			name:    "full flag set",
			failure: false,
			set:     fullSet,
		},
		{
			name:    "auth flag set",
			failure: true,
			set:     authSet,
		},
		{
			name:    "empty flag set",
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := get(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

			if test.failure {
				if err == nil {
					t.Errorf("get should have returned err")
				}

				return
			}

			if err != nil {
				t.Errorf("get returned err: %v", err)
			}
		})
	}
}
