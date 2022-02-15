// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"testing"

	"github.com/spf13/afero"
)

func TestPipeline_Config_Generate(t *testing.T) {
	// setup tests
	tests := []struct {
		config  *Config
		failure bool
	}{
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				Type:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				Type:   "go",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				Type:   "java",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				Type:   "node",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				Path:   "/tmp",
				Type:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				Stages: true,
				Type:   "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		// setup filesystem
		appFS = afero.NewMemMapFs()

		err := test.config.Generate()

		if test.failure {
			if err == nil {
				t.Errorf("Generate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Generate returned err: %v", err)
		}
	}
}
