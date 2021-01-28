// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestAction_CompletionGenerate(t *testing.T) {
	// setup flags
	bashSet := flag.NewFlagSet("test", 0)
	bashSet.Bool("bash", true, "doc")

	zshSet := flag.NewFlagSet("test", 0)
	zshSet.Bool("zsh", true, "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
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
		err := completionGenerate(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("completionGenerate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("completionGenerate returned err: %v", err)
		}
	}
}
