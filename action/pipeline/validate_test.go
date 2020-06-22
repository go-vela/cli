// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"flag"
	"testing"

	"github.com/go-vela/compiler/compiler/native"

	"github.com/urfave/cli/v2"
)

func TestPipeline_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
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
			failure: true,
			config: &Config{
				Action: "generate",
				File:   "",
				Type:   "",
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

func TestPipeline_Config_ValidateFile(t *testing.T) {
	// setup types
	c := cli.NewContext(nil, flag.NewFlagSet("test", 0), nil)

	// create a vela client
	client, err := native.New(c)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "default.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "go.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "java.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "node.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.ValidateFile(client)

		if test.failure {
			if err == nil {
				t.Errorf("ValidateFile should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("ValidateFile returned err: %v", err)
		}
	}
}
