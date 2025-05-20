// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestSecret_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup tests
	tests := []struct {
		failure bool
		cmd     *cli.Command
		args    []string
	}{
		{
			failure: false,
			cmd:     test.Command(s.URL, add, CommandAdd.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "--name", "test", "--value", "super-secret"},
		},
		{
			failure: false,
			cmd:     test.Command(s.URL, add, CommandAdd.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "--name", "test", "--value", "super-secret", "--commands", "false", "--substitution", "false"},
		},
		{
			failure: false,
			cmd:     test.Command(s.URL, add, CommandAdd.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "--file", "../../action/secret/testdata/repo.yml"},
		},
		{
			failure: false,
			cmd:     test.Command(s.URL, add, CommandAdd.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "--name", "test", "--value", "@../../action/secret/testdata/repo.yml"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, add, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

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
