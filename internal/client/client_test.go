// SPDX-License-Identifier: Apache-2.0

package client

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestClient_Parse(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	fullCmd := &cli.Command{
		Name:  "test",
		Usage: "Test command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api.addr",
				Value: s.URL,
			},
			&cli.StringFlag{
				Name:  "api.token.access",
				Value: test.TestTokenGood,
			},
			&cli.StringFlag{
				Name:  "api.token.refresh",
				Value: "superSecretRefreshToken",
			},
		},
	}

	fullTknCmd := &cli.Command{
		Name:  "test",
		Usage: "Test command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api.addr",
				Value: s.URL,
			},
			&cli.StringFlag{
				Name:  "api.token.access",
				Value: "superSecretAccessToken",
			},
			&cli.StringFlag{
				Name:  "api.token.refresh",
				Value: "superSecretRefreshToken",
			},
		},
	}

	serverCmd := &cli.Command{
		Name:  "test",
		Usage: "Test command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api.addr",
				Value: s.URL,
			},
		},
	}

	tokenCmd := &cli.Command{
		Name:  "test",
		Usage: "Test command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api.token",
				Value: "superSecretToken",
			},
		},
	}

	// setup tests
	tests := []struct {
		failure bool
		cmd     *cli.Command
	}{
		{
			failure: false,
			cmd:     fullCmd,
		},
		{
			failure: false,
			cmd:     fullTknCmd,
		},
		{
			failure: true,
			cmd:     serverCmd,
		},
		{
			failure: true,
			cmd:     tokenCmd,
		},
	}

	// run tests
	for _, test := range tests {
		_, err := Parse(test.cmd)

		if test.failure {
			if err == nil {
				t.Errorf("Parse should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Parse returned err: %v", err)
		}
	}
}

func TestClient_ParseEmptyToken(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	fullCmd := &cli.Command{
		Name:  "test",
		Usage: "Test command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api.addr",
				Value: s.URL,
			},
		},
	}

	// setup tests
	tests := []struct {
		failure bool
		cmd     *cli.Command
	}{
		{
			failure: false,
			cmd:     fullCmd,
		},
		{
			failure: true,
			cmd:     new(cli.Command),
		},
	}

	// run tests
	for _, test := range tests {
		_, err := ParseEmptyToken(test.cmd)

		if test.failure {
			if err == nil {
				t.Errorf("ParseEmptyToken should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("ParseEmptyToken returned err: %v", err)
		}
	}
}
