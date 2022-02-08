// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"flag"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/go-vela/cli/internal"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/compiler/native"
	"github.com/go-vela/server/mock/server"

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
				Action: internal.ActionCompile,
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionExpand,
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionGenerate,
				File:   ".vela.yml",
				Type:   "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				File:   "default.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:        internal.ActionValidate,
				File:          "default.yml",
				Path:          "testdata",
				Type:          "",
				TemplateFiles: []string{"nottwoelements"},
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionView,
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionGenerate,
				File:   "",
				Type:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
				Org:    "",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionView,
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
		name    string
		failure bool
		config  *Config
	}{
		{
			name:    "default",
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				File:   "default.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "go pipeline",
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				File:   "go.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "java pipeline",
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				File:   "java.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "node pipeline",
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				File:   "node.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "stages default",
			failure: false,
			config: &Config{
				Action:   internal.ActionValidate,
				File:     "default_stages_template.yml",
				Path:     "testdata",
				Type:     "",
				Template: true,
			},
		},
		{
			name:    "pipeline with template (remote)",
			failure: false,
			config: &Config{
				Action:   internal.ActionValidate,
				File:     "default_template.yml",
				Path:     "testdata",
				Type:     "",
				Template: true,
			},
		},
		{
			name:    "pipeline with template (local override)",
			failure: false,
			config: &Config{
				Action:        internal.ActionValidate,
				File:          "default_template.yml",
				Path:          "testdata",
				Type:          "",
				Template:      true,
				TemplateFiles: []string{"sample:testdata/templates/template.yml"},
			},
		},
		{
			name:    "pipeline with multiple template (local overrides)",
			failure: false,
			config: &Config{
				Action:        internal.ActionValidate,
				File:          "default_multi_template.yml",
				Path:          "testdata",
				Type:          "",
				Template:      true,
				TemplateFiles: []string{"sample:testdata/templates/template.yml", "sample2:testdata/templates/template2.yml"},
			},
		},
		{
			name:    "default without template but wants to use template",
			failure: true,
			config: &Config{
				Action:   internal.ActionValidate,
				File:     "default.yml",
				Path:     "testdata",
				Type:     "",
				Template: true,
			},
		},
		{
			name:    "pipeline with multiple template (local overrides) template mismatch",
			failure: true,
			config: &Config{
				Action:        internal.ActionValidate,
				File:          "default_multi_template.yml",
				Path:          "testdata",
				Type:          "",
				Template:      true,
				TemplateFiles: []string{"sample2:testdata/templates/template2.yml"},
			},
		},
		{
			name:    "pipeline with template (local override), wrong name in override",
			failure: true,
			config: &Config{
				Action:        internal.ActionValidate,
				File:          "default_template.yml",
				Path:          "testdata",
				Type:          "",
				Template:      true,
				TemplateFiles: []string{"foo:testdata/templates/template.yml"},
			},
		},
	}

	// run tests
	for _, test := range tests {
		isLocal := len(test.config.TemplateFiles) > 0

		err := test.config.ValidateLocal(client.WithLocal(isLocal))

		if test.failure {
			if err == nil {
				t.Errorf("(%s) ValidateLocal should have returned err", test.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("(%s) ValidateLocal returned err: %v", test.name, err)
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
				Action: internal.ActionValidate,
				Org:    "github",
				Repo:   "octocat",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				Org:    "github",
				Repo:   "octocat",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				Org:    "github",
				Repo:   "octocat",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				Org:    "github",
				Repo:   "octocat",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionValidate,
				Org:    "github",
				Repo:   "octocat",
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionValidate,
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

func Test_validateFile(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	type args struct {
		path   string
		create string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"invalid .vela.yaml", args{path: path.Join(cwd, ".vela.yaml")}, path.Join(cwd, ".vela.yaml"), true},
		{"invalid .vela.yml", args{path: path.Join(cwd, ".vela.yml")}, path.Join(cwd, ".vela.yml"), true},
		{"valid .vela.yaml", args{path: path.Join(cwd, ".vela.yaml"), create: ".vela.yaml"}, path.Join(cwd, ".vela.yaml"), false},
		{"update to .vela.yaml", args{path: path.Join(cwd, ".vela.yml"), create: ".vela.yaml"}, path.Join(cwd, ".vela.yaml"), false},
		{"valid .vela.yml", args{path: path.Join(cwd, ".vela.yml"), create: ".vela.yml"}, path.Join(cwd, ".vela.yml"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// remove existing .vela.yml and .vela.yaml
			for _, file := range []string{".vela.yml", ".vela.yaml"} {
				os.Remove(path.Join(cwd, file))
			}
			// create file if specified
			if tt.args.create != "" {
				_, err := os.Create(path.Join(cwd, tt.args.create))
				if err != nil {
					t.Error(err)
				}
			}
			got, err := validateFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
