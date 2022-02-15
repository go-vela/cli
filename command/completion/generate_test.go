// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package completion

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCompletion_Generate(t *testing.T) {
	// setup flags
	bashSet := flag.NewFlagSet("test", 0)
	bashSet.Bool("bash", true, "doc")

	zshSet := flag.NewFlagSet("test", 0)
	zshSet.Bool("zsh", true, "doc")

	// setup tests
	tests := []struct {
		set     *flag.FlagSet
		failure bool
	}{
		{
			failure: false,
			set:     bashSet,
		},
		{
			failure: false,
			set:     zshSet,
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
