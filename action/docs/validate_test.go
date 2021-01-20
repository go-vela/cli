// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docs

import (
	"testing"
)

func TestDocs_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:   "generate",
				Markdown: true,
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "generate",
				Man:    true,
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "generate",
				Markdown: true,
				Man:      true,
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "generate",
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
