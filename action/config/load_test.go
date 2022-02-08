// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"flag"
	"testing"

	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/test"
	"github.com/spf13/afero"

	"github.com/urfave/cli/v2"
)

func TestConfig_Config_Load(t *testing.T) {
	// setup app
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: internal.FlagConfig},
		&cli.StringFlag{Name: internal.FlagAPIAddress},
		&cli.StringFlag{Name: internal.FlagAPIToken},
		&cli.StringFlag{Name: internal.FlagAPIAccessToken},
		&cli.StringFlag{Name: internal.FlagAPIRefreshToken},
		&cli.StringFlag{Name: internal.FlagAPIVersion},
		&cli.StringFlag{Name: internal.FlagLogLevel},
		&cli.StringFlag{Name: internal.FlagNoGit},
		&cli.StringFlag{Name: internal.FlagOutput},
		&cli.StringFlag{Name: internal.FlagOrg},
		&cli.StringFlag{Name: internal.FlagRepo},
		&cli.StringFlag{Name: internal.FlagSecretEngine},
		&cli.StringFlag{Name: internal.FlagSecretType},
		&cli.StringFlag{Name: internal.FlagCompilerGithubDriver},
		&cli.StringFlag{Name: internal.FlagCompilerGitHubURL},
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
				Action: internal.ActionLoad,
				File:   "testdata/config.yml",
			},
			set: configSet,
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionLoad,
				File:   "testdata/config.yml",
			},
			set: fullSet,
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionLoad,
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
			Action:       internal.ActionGenerate,
			File:         test.config.File,
			Addr:         ctx.String(internal.FlagAPIAddress),
			Token:        ctx.String(internal.FlagAPIToken),
			AccessToken:  ctx.String(internal.FlagAPIAccessToken),
			RefreshToken: ctx.String(internal.FlagAPIRefreshToken),
			Version:      ctx.String(internal.FlagAPIVersion),
			LogLevel:     ctx.String(internal.FlagLogLevel),
			Engine:       ctx.String(internal.FlagSecretEngine),
			Type:         ctx.String(internal.FlagSecretType),
			GitHub: &GitHub{
				Token: ctx.String(internal.FlagCompilerGitHubToken),
				URL:   ctx.String(internal.FlagCompilerGitHubURL),
			},
			Output: ctx.String(internal.FlagOutput),
			Org:    ctx.String(internal.FlagOrg),
			Repo:   ctx.String(internal.FlagRepo),
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
