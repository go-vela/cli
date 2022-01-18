// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"

	"github.com/urfave/cli/v2"
)

func TestRepo_Sync(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	authSet := flag.NewFlagSet("test", 0)
	authSet.String("api.addr", s.URL, "doc")
	authSet.String("api.token.access", test.TestTokenGood, "doc")
	authSet.String("api.token.refresh", "superSecretRefreshToken", "doc")

	syncSet := flag.NewFlagSet("test", 0)
	syncSet.String("api.addr", s.URL, "doc")
	syncSet.String("api.token.access", test.TestTokenGood, "doc")
	syncSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	syncSet.String("org", "github", "doc")
	syncSet.String("repo", "octocat", "doc")

	syncAllSet := flag.NewFlagSet("test", 0)
	syncAllSet.String("api.addr", s.URL, "doc")
	syncAllSet.String("api.token.access", test.TestTokenGood, "doc")
	syncAllSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	syncAllSet.String("org", "github", "doc")
	syncAllSet.Bool("all", true, "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     syncSet,
		},
		{
			failure: false,
			set:     syncAllSet,
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
		err := sync(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("sync should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("sync returned err: %v", err)
		}
	}
}
