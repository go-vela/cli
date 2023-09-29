// SPDX-License-Identifier: Apache-2.0

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
