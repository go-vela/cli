// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestRepo_Update(t *testing.T) {
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
			cmd:     test.TestCommand(s.URL, update, CommandUpdate.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "--branch", "main"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, update, CommandUpdate.Flags),
			args:    []string{"--org", "github"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, update, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("update should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("update returned err: %v", err)
		}
	}
}
