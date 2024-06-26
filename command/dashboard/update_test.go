// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestDashboard_Update(t *testing.T) {
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
	fullSet.String("id", "c8da1302-07d6-11ea-882f-4893bca275b8", "doc")
	fullSet.String("add-repos", "Org-1/Repo-1,Org-2/Repo-2", "doc")
	fullSet.String("drop-repos", "Org-3/Repo-3", "doc")
	fullSet.String("target-repos", "Org-4/Repo-4", "doc")
	fullSet.String("add-admins", "octocat,octokitty", "doc")
	fullSet.String("drop-admins", "octokitten", "doc")
	fullSet.String("name", "octo-dashboard", "doc")
	fullSet.String("branches", "main,dev", "doc")
	fullSet.String("events", "push,tag", "doc")
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
		err := update(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("update should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("update returned err: %v", err)
		}
	}
}
