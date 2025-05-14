// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestPipeline_Validate(t *testing.T) {
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
			cmd:     test.Command(s.URL, validate, CommandValidate.Flags),
			args:    []string{"--file", "testdata/.vela.yml"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, validate, CommandValidate.Flags),
			args:    []string{"--file", "empty.yml"},
		},
		{
			failure: true,
			cmd:     test.Command(s.URL, validate, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("validate returned err: %v", err)
		}
	}
}
