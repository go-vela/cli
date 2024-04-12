// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"flag"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestWorker_Get(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token.access", test.TestTokenGood, "doc")
	fullSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	fullSet.String("output", "json", "doc")
	fullSet.Int64("before", 42, "doc")
	fullSet.Int64("after", 0, "doc")
	fullSet.String("active", "true", "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     fullSet,
		},
		{
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := get(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

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
