// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestConfig_Remove(t *testing.T) {
	// setup flags
	set := flag.NewFlagSet("test", 0)
	set.String("config", "../../action/config/testdata/remove.yml", "doc")
	set.Bool("api.addr", true, "doc")
	set.Bool("api.token", true, "doc")
	set.Bool("api.version", true, "doc")
	set.Bool("log.level", true, "doc")
	set.Bool("no-git", true, "doc")
	set.Bool("output", true, "doc")
	set.Bool("org", true, "doc")
	set.Bool("repo", true, "doc")
	set.Bool("secret.engine", true, "doc")
	set.Bool("secret.type", true, "doc")

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
		err := remove(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("remove should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("remove returned err: %v", err)
		}
	}
}
