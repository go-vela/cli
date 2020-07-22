// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestAction_ConfigUpdate(t *testing.T) {
	// setup flags
	set := flag.NewFlagSet("test", 0)
	set.String("config", "config/testdata/generate.yml", "doc")
	set.String("api.addr", "https://vela-server.localhost", "doc")
	set.String("api.token", "superSecretToken", "doc")
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
		err := configUpdate(cli.NewContext(nil, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("configUpdate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("configUpdate returned err: %v", err)
		}
	}
}
