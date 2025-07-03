// SPDX-License-Identifier: Apache-2.0

package build

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestBuild_Restart(t *testing.T) {
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
			cmd:     test.Command(s.URL, restart, CommandRestart.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "--build", "1"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, approve, CommandRestart.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "cat"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, restart, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("restart should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("restart returned err: %v", err)
		}
	}
}
