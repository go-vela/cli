// SPDX-License-Identifier: Apache-2.0

package version

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestVersion_Version(t *testing.T) {
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
			cmd:     test.TestCommand(s.URL, runVersion, CommandVersion.Flags),
			args:    []string{"--output", "spew"},
		},
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, runVersion, CommandVersion.Flags),
			args:    []string{"--output", "dump"},
		},
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, runVersion, CommandVersion.Flags),
			args:    []string{"--output", "json"},
		},
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, runVersion, CommandVersion.Flags),
			args:    []string{"--output", "yaml"},
		},
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, runVersion, CommandVersion.Flags),
			args:    []string{},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("version should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("version returned err: %v", err)
		}
	}
}
