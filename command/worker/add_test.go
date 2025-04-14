// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/worker/mock/worker"
)

func TestWorker_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	w := httptest.NewServer(worker.FakeHandler())

	// setup tests
	tests := []struct {
		failure bool
		cmd     *cli.Command
		args    []string
	}{
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, add, CommandAdd.Flags),
			args:    []string{"--wa", w.URL, "--wh", "worker"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, add, CommandAdd.Flags),
			args:    []string{"--wh", "worker"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, add, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("add should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("add returned err: %v", err)
		}
	}
}
