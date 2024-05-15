// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestSettings_View(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token.access", test.TestTokenGood, "doc")
	fullSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
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
	}

	// run tests
	for _, test := range tests {
		err := view(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("view should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("view returned err: %v", err)
		}
	}
}
