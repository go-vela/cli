// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package completion

import (
	"testing"

	"github.com/go-vela/cli/internal"
)

func TestCompletion_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: internal.ActionGenerate,
				Bash:   true,
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionGenerate,
				Zsh:    true,
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionGenerate,
				Bash:   true,
				Zsh:    true,
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionGenerate,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}
