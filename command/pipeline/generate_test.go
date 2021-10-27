// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestPipeline_Generate(t *testing.T) {
	// setup flags
	set := flag.NewFlagSet("test", 0)
	set.String("file", "generate.yml", "doc")
	set.String("path", "../../action/pipeline/testdata", "doc")

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
