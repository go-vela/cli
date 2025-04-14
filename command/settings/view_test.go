// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestSettings_View(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup tests
	tests := []struct {
		failure bool
		cmd     *cli.Command
	}{
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, view, CommandView.Flags),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), []string{"test"})

		if test.failure {
			if err == nil {
				t.Errorf("view should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("view returned err: %v", err)
		}
	}
}
