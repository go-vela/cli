// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestPipeline_Validate(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	authSet := flag.NewFlagSet("test", 0)
	authSet.String("api.addr", s.URL, "doc")
	authSet.String("api.token.access", test.TestTokenGood, "doc")
	authSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	authSet.String("clone-image", "target/vela-git:latest", "doc")

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token.access", test.TestTokenGood, "doc")
	fullSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	fullSet.String("org", "github", "doc")
	fullSet.String("repo", "octocat", "doc")
	fullSet.String("output", "json", "doc")
	fullSet.String("pipeline-type", "yaml", "doc")
	fullSet.String("branch", "main", "doc")
	fullSet.String("comment", "comment", "doc")
	fullSet.String("event", "push", "doc")
	fullSet.String("status", "success", "doc")
	fullSet.String("tag", "v0.0.0", "doc")
	fullSet.String("target", "production", "doc")
	fullSet.String("file-changeset", "README.md,main,go", "doc")
	fullSet.Uint64("compiler-starlark-exec-limit", 10000, "doc")
	fullSet.Bool("remote", true, "doc")
	fullSet.String("clone-image", "target/vela-git:latest", "doc")

	// setup flags
	localSet := flag.NewFlagSet("test", 0)
	localSet.String("file", "generate.yml", "doc")
	localSet.String("path", "../../action/pipeline/testdata", "doc")
	localSet.String("clone-image", "target/vela-git:latest", "doc")

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
