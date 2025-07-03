// SPDX-License-Identifier: Apache-2.0

package action

import (
	"net/http/httptest"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/mock/server"
)

func TestAction_Load(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	authCmd := &cli.Command{
		Name: "auth",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config/testdata/empty.yml",
			},
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

	fullCmd := &cli.Command{
		Name: "full",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config/testdata/empty.yml",
			},
			&cli.StringFlag{
				Name:  "api.addr",
				Value: "https://vela-server.localhost",
			},
			&cli.StringFlag{
				Name:  "api.token.access",
				Value: test.TestTokenGood,
			},
			&cli.StringFlag{
				Name:  "api.token.refresh",
				Value: "superSecretRefreshToken",
			},
			&cli.StringFlag{
				Name:  "api.version",
				Value: "1",
			},
			&cli.StringFlag{
				Name:  "log.level",
				Value: "info",
			},
			&cli.StringFlag{
				Name:  "no-git",
				Value: "true",
			},
			&cli.StringFlag{
				Name:  "output",
				Value: "json",
			},
			&cli.StringFlag{
				Name:  "org",
				Value: "github",
			},
			&cli.StringFlag{
				Name:  "repo",
				Value: "octocat",
			},
			&cli.StringFlag{
				Name:  "secret.engine",
				Value: "native",
			},
			&cli.StringFlag{
				Name:  "secret.type",
				Value: "repo",
			},
			&cli.StringFlag{
				Name:  "compiler.github.driver",
				Value: "true",
			},
			&cli.StringFlag{
				Name:  "compiler.github.url",
				Value: "github.com",
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
			cmd:     authCmd,
		},
		{
			failure: false,
			cmd:     new(cli.Command),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), []string{test.cmd.Name})
		if err != nil {
			t.Errorf("Run returned err: %v", err)
		}

		err = Load(test.cmd)

		if test.failure {
			if err == nil {
				t.Errorf("Load should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Load returned err: %v", err)
		}
	}
}
