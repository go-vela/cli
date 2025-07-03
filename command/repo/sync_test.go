// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestRepo_Sync(t *testing.T) {
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
			cmd:     test.Command(s.URL, sync, CommandSync.Flags),
			args:    []string{"--org", "github", "--repo", "octocat"},
		},
		{
			failure: false,
			cmd:     test.Command(s.URL, sync, CommandSync.Flags),
			args:    []string{"--org", "github", "--all"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, sync, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("sync should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("sync returned err: %v", err)
		}
	}
}
