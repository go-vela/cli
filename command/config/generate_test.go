// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestConfig_Generate(t *testing.T) {
	// setup flags
	set := flag.NewFlagSet("test", 0)
	set.String("config", "../../action/config/testdata/generate.yml", "doc")
	set.String("api.addr", "https://vela-server.localhost", "doc")
	set.String("api.token", "superSecretToken", "doc")
	set.String("api.token.access", "superSecretAccessToken", "doc")
	set.String("api.token.refresh", "superSecretRefreshToken", "doc")
	set.String("api.version", "1", "doc")
	set.String("log.level", "info", "doc")
	set.String("output", "json", "doc")
	set.String("org", "github", "doc")
	set.String("repo", "octocat", "doc")
	set.String("secret.engine", "native", "doc")
	set.String("secret.type", "repo", "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     set,
		},
		{
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := generate(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("generate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("generate returned err: %v", err)
		}
	}
}
