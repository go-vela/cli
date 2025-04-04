// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/worker/mock/worker"
)

func TestWorker_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create mock worker server
	w := httptest.NewServer(worker.FakeHandler())

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

	noAddressSet := flag.NewFlagSet("test", 0)
	noAddressSet.String("api.addr", s.URL, "doc")
	noAddressSet.String("api.token.access", test.TestTokenGood, "doc")
	noAddressSet.String("api.token.refresh", "superSecretRefreshToken", "doc")

	noHostnameSet := flag.NewFlagSet("test", 0)
	noHostnameSet.String("api.addr", s.URL, "doc")
	noHostnameSet.String("api.token.access", test.TestTokenGood, "doc")
	noHostnameSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	noHostnameSet.String("worker.address", w.URL, "doc")

	badAddressSet := flag.NewFlagSet("test", 0)
	badAddressSet.String("api.addr", s.URL, "doc")
	badAddressSet.String("api.token.access", test.TestTokenGood, "doc")
	badAddressSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	badAddressSet.String("worker.address", "::", "doc")

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
		{
			failure: true,
			set:     noAddressSet,
		},
		{
			failure: false,
			set:     noHostnameSet,
		},
		{
			failure: true,
			set:     badAddressSet,
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
