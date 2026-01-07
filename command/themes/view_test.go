// SPDX-License-Identifier: Apache-2.0

package themes

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestThemes_View(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup tests
	tests := []struct {
		name    string
		failure bool
		cmd     *cli.Command
		args    []string
	}{
		{
			name:    "default output",
			failure: false,
			cmd:     test.Command(s.URL, view, CommandView.Flags),
			args:    []string{},
		},
		{
			name:    "json output",
			failure: false,
			cmd:     test.Command(s.URL, view, CommandView.Flags),
			args:    []string{"--output", "json"},
		},
		{
			name:    "yaml output",
			failure: false,
			cmd:     test.Command(s.URL, view, CommandView.Flags),
			args:    []string{"--output", "yaml"},
		},
		{
			name:    "dump output",
			failure: false,
			cmd:     test.Command(s.URL, view, CommandView.Flags),
			args:    []string{"--output", "dump"},
		},
		{
			name:    "spew output",
			failure: false,
			cmd:     test.Command(s.URL, view, CommandView.Flags),
			args:    []string{"--output", "spew"},
		},
		{
			name:    "short flag alias",
			failure: false,
			cmd:     test.Command(s.URL, view, CommandView.Flags),
			args:    []string{"--op", "json"},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

			if test.failure {
				if err == nil {
					t.Errorf("view should have returned err")
				}

				return
			}

			if err != nil {
				t.Errorf("view returned err: %v", err)
			}
		})
	}
}
