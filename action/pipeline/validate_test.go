// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/compiler/compiler/native"
	"github.com/go-vela/mock/server"
	"github.com/go-vela/sdk-go/vela"

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
				Action: "compile",
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "expand",
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
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
				Action: "validate",
				File:   "default.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Output: "",
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
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "",
				Output: "",
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

func TestPipeline_Config_ValidateLocal(t *testing.T) {
	// setup types
	c := cli.NewContext(&cli.App{Name: "vela", Version: "v0.0.0"}, flag.NewFlagSet("test", 0), nil)

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
		err := test.config.ValidateLocal(client)

		if test.failure {
			if err == nil {
				t.Errorf("ValidateLocal should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("ValidateLocal returned err: %v", err)
		}
	}
}

func TestPipeline_Config_ValidateRemote(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
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
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "not-found",
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.ValidateRemote(client)

		if test.failure {
			if err == nil {
				t.Errorf("ValidateRemote should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("ValidateRemote returned err: %v", err)
		}
	}
}
