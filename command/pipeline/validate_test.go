// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/mock/server"
	"github.com/urfave/cli/v2"
)

func TestPipeline_Validate(t *testing.T) {
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
	fullSet.String("output", "json", "doc")
	fullSet.String("pipeline-type", "yaml", "doc")
	fullSet.Bool("remote", true, "doc")

	// setup flags
	localSet := flag.NewFlagSet("test", 0)
	localSet.String("file", "generate.yml", "doc")
	localSet.String("path", "../../action/pipeline/testdata", "doc")

	// setup tests
	tests := []struct {
		name    string
		failure bool
		set     *flag.FlagSet
	}{
		{
			name:    "fullSet",
			failure: false,
			set:     fullSet,
		},
		{
			name:    "localSet",
			failure: false,
			set:     localSet,
		},
		{
			name:    "authSet",
			failure: true,
			set:     authSet,
		},
		{
			name:    "newFlagSet",
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := validate(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("(%s) validate should have returned err", test.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("(%s) validate returned err: %v", test.name, err)
		}
	}
}
