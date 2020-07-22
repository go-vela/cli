// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"flag"
	"testing"

	"github.com/spf13/afero"

	"github.com/urfave/cli/v2"
)

func TestConfig_Config_Load(t *testing.T) {
	// setup flags
	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", "https://vela-server.localhost", "doc")
	fullSet.String("api.token", "superSecretToken", "doc")
	fullSet.String("api.version", "1", "doc")
	fullSet.String("log.level", "info", "doc")
	fullSet.String("output", "json", "doc")
	fullSet.String("org", "github", "doc")
	fullSet.String("repo", "octocat", "doc")
	fullSet.String("secret.engine", "native", "doc")
	fullSet.String("secret.type", "repo", "doc")

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
		ctx := cli.NewContext(nil, test.set, nil)

		// setup filesystem
		appFS = afero.NewMemMapFs()

		// create test config for generating file
		config := &Config{
			Action:   "generate",
			File:     test.config.File,
			Addr:     ctx.String("api.addr"),
			Token:    ctx.String("api.token"),
			Version:  ctx.String("api.version"),
			LogLevel: ctx.String("log.level"),
			Engine:   ctx.String("secret.engine"),
			Type:     ctx.String("secret.type"),
			Output:   ctx.String("output"),
			Org:      ctx.String("org"),
			Repo:     ctx.String("repo"),
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
