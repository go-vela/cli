// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestAction_ConfigView(t *testing.T) {
	// setup flags
	set := flag.NewFlagSet("test", 0)
	set.String("file", "config/testdata/config.yml", "doc")

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
		err := configView(cli.NewContext(nil, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("configView should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("configView returned err: %v", err)
		}
	}
}
