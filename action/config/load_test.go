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
	set := flag.NewFlagSet("test", 0)
	set.String("api.addr", "https://vela-server.localhost", "doc")
	set.String("api.token", "superSecretToken", "doc")
	set.String("api.version", "1", "doc")
	set.String("log.level", "info", "doc")
	set.String("output", "json", "doc")
	set.String("org", "github", "doc")
	set.String("repo", "octocat", "doc")
	set.String("secret.engine", "native", "doc")
	set.String("type", "repo", "doc")

	addrSet := flag.NewFlagSet("test", 0)
	addrSet.String("api.token", "superSecretToken", "doc")
	addrSet.String("api.version", "1", "doc")
	addrSet.String("log.level", "info", "doc")
	addrSet.String("output", "json", "doc")
	addrSet.String("org", "github", "doc")
	addrSet.String("repo", "octocat", "doc")
	addrSet.String("secret.engine", "native", "doc")
	addrSet.String("type", "repo", "doc")

	tokenSet := flag.NewFlagSet("test", 0)
	tokenSet.String("api.addr", "https://vela-server.localhost", "doc")
	tokenSet.String("api.version", "1", "doc")
	tokenSet.String("log.level", "info", "doc")
	tokenSet.String("output", "json", "doc")
	tokenSet.String("org", "github", "doc")
	tokenSet.String("repo", "octocat", "doc")
	tokenSet.String("secret.engine", "native", "doc")
	tokenSet.String("type", "repo", "doc")

	versionSet := flag.NewFlagSet("test", 0)
	versionSet.String("api.addr", "https://vela-server.localhost", "doc")
	versionSet.String("api.token", "superSecretToken", "doc")
	versionSet.String("log.level", "info", "doc")
	versionSet.String("output", "json", "doc")
	versionSet.String("org", "github", "doc")
	versionSet.String("repo", "octocat", "doc")
	versionSet.String("secret.engine", "native", "doc")
	versionSet.String("type", "repo", "doc")

	logSet := flag.NewFlagSet("test", 0)
	logSet.String("api.addr", "https://vela-server.localhost", "doc")
	logSet.String("api.token", "superSecretToken", "doc")
	logSet.String("api.version", "1", "doc")
	logSet.String("output", "json", "doc")
	logSet.String("org", "github", "doc")
	logSet.String("repo", "octocat", "doc")
	logSet.String("secret.engine", "native", "doc")
	logSet.String("type", "repo", "doc")

	outputSet := flag.NewFlagSet("test", 0)
	outputSet.String("api.addr", "https://vela-server.localhost", "doc")
	outputSet.String("api.token", "superSecretToken", "doc")
	outputSet.String("api.version", "1", "doc")
	outputSet.String("log.level", "info", "doc")
	outputSet.String("org", "github", "doc")
	outputSet.String("repo", "octocat", "doc")
	outputSet.String("secret.engine", "native", "doc")
	outputSet.String("type", "repo", "doc")

	orgSet := flag.NewFlagSet("test", 0)
	orgSet.String("api.addr", "https://vela-server.localhost", "doc")
	orgSet.String("api.token", "superSecretToken", "doc")
	orgSet.String("api.version", "1", "doc")
	orgSet.String("log.level", "info", "doc")
	orgSet.String("output", "json", "doc")
	orgSet.String("repo", "octocat", "doc")
	orgSet.String("secret.engine", "native", "doc")
	orgSet.String("type", "repo", "doc")

	repoSet := flag.NewFlagSet("test", 0)
	repoSet.String("api.addr", "https://vela-server.localhost", "doc")
	repoSet.String("api.token", "superSecretToken", "doc")
	repoSet.String("api.version", "1", "doc")
	repoSet.String("log.level", "info", "doc")
	repoSet.String("output", "json", "doc")
	repoSet.String("org", "github", "doc")
	repoSet.String("secret.engine", "native", "doc")
	repoSet.String("type", "repo", "doc")

	engineSet := flag.NewFlagSet("test", 0)
	engineSet.String("api.addr", "https://vela-server.localhost", "doc")
	engineSet.String("api.token", "superSecretToken", "doc")
	engineSet.String("api.version", "1", "doc")
	engineSet.String("log.level", "info", "doc")
	engineSet.String("output", "json", "doc")
	engineSet.String("org", "github", "doc")
	engineSet.String("repo", "octocat", "doc")
	engineSet.String("type", "repo", "doc")

	typeSet := flag.NewFlagSet("test", 0)
	typeSet.String("api.addr", "https://vela-server.localhost", "doc")
	typeSet.String("api.token", "superSecretToken", "doc")
	typeSet.String("api.version", "1", "doc")
	typeSet.String("log.level", "info", "doc")
	typeSet.String("output", "json", "doc")
	typeSet.String("org", "github", "doc")
	typeSet.String("repo", "octocat", "doc")
	typeSet.String("secret.engine", "native", "doc")

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
			set: set,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: addrSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: tokenSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: versionSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: logSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: outputSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: orgSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: repoSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: engineSet,
		},
		{
			failure: true,
			config: &Config{
				Action: "load",
				File:   "testdata/config.yml",
			},
			set: typeSet,
		},
	}

	// run tests
	for _, test := range tests {
		// setup filesystem
		appFS = afero.NewMemMapFs()

		// create test config for generating file
		config := &Config{
			Action: "generate",
			File:   test.config.File,
		}

		// generate config file
		err := config.Generate()
		if err != nil {
			t.Errorf("unable to generate config: %v", err)
		}

		err = test.config.Load(cli.NewContext(nil, test.set, nil))

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
