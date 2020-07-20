// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/urfave/cli/v2"
)

func TestAction_RepoGet(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("token", "superSecretToken", "doc")
	fullSet.Int("page", 1, "doc")
	fullSet.Int("per.page", 10, "doc")
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
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := repoGet(cli.NewContext(nil, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("repoGet should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("repoGet returned err: %v", err)
		}
	}
}
