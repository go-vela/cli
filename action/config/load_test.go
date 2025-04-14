// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
)

func TestConfig_Config_Load(t *testing.T) {
	// setup app
	fullFlags := []cli.Flag{
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
	}

	// setup tests
	tests := []struct {
		failure bool
		config  *Config
		args    []string
		flags   []cli.Flag
	}{
		{
			failure: false,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			args: []string{"test", "config"},
		},
		{
			failure: false,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			args:  []string{"test", "repo"},
			flags: fullFlags,
		},
	}

	// run tests
	for _, test := range tests {
		// setup filesystem
		appFS = afero.NewMemMapFs()

		cmd := new(cli.Command)
		cmd.Name = "test"
		cmd.Commands = []*cli.Command{
			{
				Name:  "config",
				Usage: "test config command",
			},
			{
				Name:  "repo",
				Usage: "test repo command",
			},
		}
		cmd.Flags = test.flags

		// create test config for generating file
		config := &Config{
			Action:       "generate",
			File:         test.config.File,
			Addr:         cmd.String("api.addr"),
			Token:        cmd.String("api.token"),
			AccessToken:  cmd.String("api.token.access"),
			RefreshToken: cmd.String("api.token.refresh"),
			Version:      cmd.String("api.version"),
			LogLevel:     cmd.String("log.level"),
			Engine:       cmd.String("secret.engine"),
			Type:         cmd.String("secret.type"),
			GitHub: &GitHub{
				Token: cmd.String("compiler.github.token"),
				URL:   cmd.String("compiler.github.url"),
			},
			Output: cmd.String("output"),
			Org:    cmd.String("org"),
			Repo:   cmd.String("repo"),
		}

		// generate config file
		err := config.Generate()
		if err != nil {
			t.Errorf("unable to generate config: %v", err)
		}

		err = cmd.Run(t.Context(), test.args)
		if err != nil {
			t.Errorf("unable to run command: %v", err)
		}

		err = test.config.Load(cmd)

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
