// SPDX-License-Identifier: Apache-2.0

package build

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestBuild_Cancel(t *testing.T) {
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
			cmd:     test.Command(s.URL, cancel, CommandCancel.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "--build", "1"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, approve, CommandCancel.Flags),
			args:    []string{"--org", "github", "--repo", "octocat", "cat"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, cancel, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("cancel should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("cancel returned err: %v", err)
		}
	}
}
