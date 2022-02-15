// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docs

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestDocs_Generate(t *testing.T) {
	// setup flags
	markdownSet := flag.NewFlagSet("test", 0)
	markdownSet.Bool("markdown", true, "doc")

	manSet := flag.NewFlagSet("test", 0)
	manSet.Bool("man", true, "doc")

	// setup tests
	tests := []struct {
		set     *flag.FlagSet
		failure bool
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
