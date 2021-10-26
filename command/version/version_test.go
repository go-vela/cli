// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package version

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestVersion_RunVersion(t *testing.T) {
	// setup flags
	dumpSet := flag.NewFlagSet("test", 0)
	dumpSet.String("output", "dump", "doc")

	jsonSet := flag.NewFlagSet("test", 0)
	jsonSet.String("output", "json", "doc")

	spewSet := flag.NewFlagSet("test", 0)
	spewSet.String("output", "spew", "doc")

	yamlSet := flag.NewFlagSet("test", 0)
	yamlSet.String("output", "yaml", "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     dumpSet,
		},
		{
			failure: false,
			set:     jsonSet,
		},
		{
			failure: false,
			set:     spewSet,
		},
		{
			failure: false,
			set:     yamlSet,
		},
		{
			failure: false,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := runVersion(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("runVersion should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("runVersion returned err: %v", err)
		}
	}
}
