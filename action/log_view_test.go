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

func TestAction_LogView(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	authSet := flag.NewFlagSet("test", 0)
	authSet.String("api.addr", s.URL, "doc")
	authSet.String("token", "superSecretToken", "doc")

	serviceSet := flag.NewFlagSet("test", 0)
	serviceSet.String("api.addr", s.URL, "doc")
	serviceSet.String("token", "superSecretToken", "doc")
	serviceSet.String("org", "github", "doc")
	serviceSet.String("repo", "octocat", "doc")
	serviceSet.Int("build", 1, "doc")
	serviceSet.Int("service", 1, "doc")
	serviceSet.String("output", "json", "doc")

	stepSet := flag.NewFlagSet("test", 0)
	stepSet.String("api.addr", s.URL, "doc")
	stepSet.String("token", "superSecretToken", "doc")
	stepSet.String("org", "github", "doc")
	stepSet.String("repo", "octocat", "doc")
	stepSet.Int("build", 1, "doc")
	stepSet.Int("step", 1, "doc")
	stepSet.String("output", "json", "doc")

	buildSet := flag.NewFlagSet("test", 0)
	buildSet.String("api.addr", s.URL, "doc")
	buildSet.String("token", "superSecretToken", "doc")
	buildSet.String("org", "github", "doc")
	buildSet.String("repo", "octocat", "doc")
	buildSet.Int("build", 1, "doc")
	buildSet.String("output", "json", "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     serviceSet,
		},
		{
			failure: false,
			set:     stepSet,
		},
		{
			failure: false,
			set:     buildSet,
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
		err := logView(cli.NewContext(nil, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("logView should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("logView returned err: %v", err)
		}
	}
}
