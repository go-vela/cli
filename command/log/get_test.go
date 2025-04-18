// SPDX-License-Identifier: Apache-2.0

package log

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestLog_Get(t *testing.T) {
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
			cmd:     test.TestCommand(s.URL, get, CommandGet.Flags),
			args:    []string{"--org", "Org-1", "--repo", "Repo-1", "--build", "1"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, get, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("get should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("get returned err: %v", err)
		}
	}
}
