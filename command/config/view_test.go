// SPDX-License-Identifier: Apache-2.0

package config

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestConfig_View(t *testing.T) {
	// setup flags
	set := flag.NewFlagSet("test", 0)
	set.String("config", "../../action/config/testdata/config.yml", "doc")

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
		err := view(cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("view should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("view returned err: %v", err)
		}
	}
}
