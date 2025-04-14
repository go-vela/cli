// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestWorker_Get(t *testing.T) {
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
			args:    []string{"--before", "16000000", "--active", "true"},
		},
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, get, CommandGet.Flags),
			args:    []string{},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, get, CommandGet.Flags),
			args:    []string{"--before", "today"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, get, CommandGet.Flags),
			args:    []string{"--after", "today"},
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

func TestWorker_parseUnix(t *testing.T) {
	tests := []struct {
		input   string
		want    int64
		wantErr bool
	}{
		{
			input: "2019-11-06T20:58:00",
			want:  1573073880,
		},
		{
			input: "10m",
			want:  time.Now().Add(-10 * time.Minute).Unix(),
		},
		{
			input: "42",
			want:  42,
		},
		{
			input:   "invalid",
			want:    0,
			wantErr: true,
		},
	}

	for _, test := range tests {
		got, err := parseUnix(test.input)

		if test.wantErr && err == nil {
			t.Errorf("parseUnix should have returned error")
		}

		if got != test.want {
			t.Errorf("parseUnix returned %d, want %d", got, test.want)
		}
	}
}
