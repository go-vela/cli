// SPDX-License-Identifier: Apache-2.0

package client

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestClient_Parse(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a simple sample app
	a := new(cli.App)
	a.Name = "vela"
	a.Version = "v1.0.0"

	// setup flags
	serverSet := flag.NewFlagSet("test", 0)
	serverSet.String("api.addr", s.URL, "doc")

	tokenSet := flag.NewFlagSet("test", 0)
	tokenSet.String("api.token", "superSecretToken", "doc")

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token.access", test.TestTokenGood, "doc")
	fullSet.String("api.token.refresh", "superSecretRefreshToken", "doc")

	fullSetTokenSet := flag.NewFlagSet("test", 0)
	fullSetTokenSet.String("api.addr", s.URL, "doc")
	fullSetTokenSet.String("api.token.access", "superSecretAccessToken", "doc")
	fullSetTokenSet.String("api.token.refresh", "superSecretRefreshToken", "doc")

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
			failure: false,
			set:     fullSetTokenSet,
		},
		{
			failure: true,
			set:     serverSet,
		},
		{
			failure: true,
			set:     tokenSet,
		},
	}

	// run tests
	for _, test := range tests {
		_, err := Parse(cli.NewContext(a, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("Parse should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Parse returned err: %v", err)
		}
	}
}

func TestClient_ParseEmptyToken(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")

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
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		_, err := ParseEmptyToken(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("ParseEmptyToken should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("ParseEmptyToken returned err: %v", err)
		}
	}
}
