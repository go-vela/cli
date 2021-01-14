// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestAction_docsGenerate(t *testing.T) {
	// setup flags
	markdownSet := flag.NewFlagSet("test", 0)
	markdownSet.Bool("markdown", true, "doc")

	manSet := flag.NewFlagSet("test", 0)
	manSet.Bool("man", true, "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     markdownSet,
		},
		{
			failure: false,
			set:     manSet,
		},
		{
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := docsGenerate(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("docsGenerate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("docsGenerate returned err: %v", err)
		}
	}
}
