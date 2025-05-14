// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestDocs_Generate(t *testing.T) {
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
			cmd:     test.Command(s.URL, generate, CommandGenerate.Flags),
			args:    []string{"--markdown", "true"},
		},
		{
			failure: false,
			cmd:     test.Command(s.URL, generate, CommandGenerate.Flags),
			args:    []string{"--man", "true"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, generate, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("generate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("generate returned err: %v", err)
		}
	}
}
