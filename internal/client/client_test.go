// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package client

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/urfave/cli/v2"
)

func TestClient_Parse(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	serverSet := flag.NewFlagSet("test", 0)
	serverSet.String("api.addr", s.URL, "doc")

	tokenSet := flag.NewFlagSet("test", 0)
	tokenSet.String("api.token", "superSecretToken", "doc")

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token", "superSecretToken", "doc")

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
			set:     serverSet,
		},
		{
			failure: true,
			set:     tokenSet,
		},
	}

	// run tests
	for _, test := range tests {
		_, err := Parse(cli.NewContext(nil, test.set, nil))

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
		_, err := ParseEmptyToken(cli.NewContext(nil, test.set, nil))

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
