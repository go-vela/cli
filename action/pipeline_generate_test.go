// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestAction_PipelineGenerate(t *testing.T) {
	// setup flags
	set := flag.NewFlagSet("test", 0)
	set.String("file", "generate.yml", "doc")
	set.String("path", "pipeline/testdata", "doc")

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
		err := pipelineGenerate(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("pipelineGenerate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("pipelineGenerate returned err: %v", err)
		}
	}
}
