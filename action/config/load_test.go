// SPDX-License-Identifier: Apache-2.0

package config

import (
	"flag"
	"testing"

	"github.com/spf13/afero"
	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/test"
)

func TestConfig_Config_Load(t *testing.T) {
	// setup app
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "config"},
		&cli.StringFlag{Name: "api.addr"},
		&cli.StringFlag{Name: "api.token"},
		&cli.StringFlag{Name: "api.token.access"},
		&cli.StringFlag{Name: "api.token.refresh"},
		&cli.StringFlag{Name: "api.version"},
		&cli.StringFlag{Name: "log.level"},
		&cli.StringFlag{Name: "no-git"},
		&cli.StringFlag{Name: "output"},
		&cli.StringFlag{Name: "org"},
		&cli.StringFlag{Name: "repo"},
		&cli.StringFlag{Name: "secret.engine"},
		&cli.StringFlag{Name: "secret.type"},
		&cli.StringFlag{Name: "compiler.github.driver"},
		&cli.StringFlag{Name: "compiler.github.url"},
	}

	// setup flags
	configSet := flag.NewFlagSet("test", 0)
	err := configSet.Parse([]string{"view", "config"})

	if err != nil {
		t.Errorf("unable to parse configset: %v", err)
	}

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", "https://vela-server.localhost", "doc")
	fullSet.String("api.token.access", test.TestTokenGood, "doc")
	fullSet.String("api.token.refresh", "superSecretRefreshToken", "doc")
	fullSet.String("api.version", "1", "doc")
	fullSet.String("log.level", "info", "doc")
	fullSet.String("no-git", "true", "doc")
	fullSet.String("output", "json", "doc")
	fullSet.String("org", "github", "doc")
	fullSet.String("repo", "octocat", "doc")
	fullSet.String("secret.engine", "native", "doc")
	fullSet.String("secret.type", "repo", "doc")
	fullSet.String("compiler.github.driver", "true", "doc")
	fullSet.String("compiler.github.url", "github.com", "doc")

	// setup tests
	tests := []struct {
		failure bool
		config  *Config
		set     *flag.FlagSet
	}{
		{
			failure: false,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: configSet,
		},
		{
			failure: false,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: fullSet,
		},
		{
			failure: false,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		// setup context
		ctx := cli.NewContext(app, test.set, nil)

		// setup filesystem
		appFS = afero.NewMemMapFs()

		// create test config for generating file
		config := &Config{
			Action:       "generate",
			File:         test.config.File,
			Addr:         ctx.String("api.addr"),
			Token:        ctx.String("api.token"),
			AccessToken:  ctx.String("api.token.access"),
			RefreshToken: ctx.String("api.token.refresh"),
			Version:      ctx.String("api.version"),
			LogLevel:     ctx.String("log.level"),
			Engine:       ctx.String("secret.engine"),
			Type:         ctx.String("secret.type"),
			GitHub: &GitHub{
				Token: ctx.String("compiler.github.token"),
				URL:   ctx.String("compiler.github.url"),
			},
			Output: ctx.String("output"),
			Org:    ctx.String("org"),
			Repo:   ctx.String("repo"),
		}

		// generate config file
		err := config.Generate()
		if err != nil {
			t.Errorf("unable to generate config: %v", err)
		}

		err = test.config.Load(ctx)

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
