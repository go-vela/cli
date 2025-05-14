// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestPipeline_Expand(t *testing.T) {
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
			cmd:     test.Command(s.URL, expand, CommandExpand.Flags),
			args:    []string{"--org", "Org-1", "--repo", "Repo-1", "--ref", "refs/heads/main"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, expand, CommandExpand.Flags),
			args:    []string{"--org", "Org-1"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, expand, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("expand should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("expand returned err: %v", err)
		}
	}
}
