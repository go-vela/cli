// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/compiler/native"
	"github.com/go-vela/server/mock/server"
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
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "exec",
				Org:    "github",
				Repo:   "octocat",
				Event:  "tag",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "exec",
				Org:    "github",
				Repo:   "octocat",
				Event:  "tag",
				Tag:    "v1.0.0",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "expand",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
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
				Action:  "get",
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "",
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
			failure: true,
			config: &Config{
				Action:        "validate",
				File:          "default.yml",
				Path:          "testdata",
				Type:          "",
				TemplateFiles: []string{"nottwoelements"},
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
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
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "",
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
	cmd := new(cli.Command)
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "clone-image",
			Value: "target/vela-git:latest",
		},
	}

	// create a vela client
	client, err := native.FromCLICommand(t.Context(), cmd)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	client.SetTemplateDepth(1)

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
				Action: "validate",
				File:   "default.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "go pipeline",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "go.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "java pipeline",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "java.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "node pipeline",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "node.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "stages default",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "default_stages_template.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "pipeline with template (remote)",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "default_template.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "pipeline with template (local override)",
			failure: false,
			config: &Config{
				Action:        "validate",
				File:          "default_template.yml",
				Path:          "testdata",
				Type:          "",
				TemplateFiles: []string{"sample:testdata/templates/template.yml"},
			},
		},
		{
			name:    "pipeline with multiple template (local overrides)",
			failure: false,
			config: &Config{
				Action:        "validate",
				File:          "default_multi_template.yml",
				Path:          "testdata",
				Type:          "",
				TemplateFiles: []string{"sample:testdata/templates/template.yml", "sample2:testdata/templates/template2.yml"},
			},
		},
		{
			name: "pipeline with multiple template (local overrides) only one template specified",
			config: &Config{
				Action:        "validate",
				File:          "default_multi_template.yml",
				Path:          "testdata",
				Type:          "",
				TemplateFiles: []string{"sample2:testdata/templates/template2.yml"},
			},
		},
		{
			name:    "pipeline with template (local override), wrong name in override",
			failure: true,
			config: &Config{
				Action:        "validate",
				File:          "default_template.yml",
				Path:          "testdata",
				Type:          "",
				TemplateFiles: []string{"foo:testdata/templates/template.yml"},
			},
		},
		{
			name:    "pipeline with rulesets - no ruledata provided",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "ruleset.yml",
				Path:   "testdata",
				Type:   "",
			},
		},
		{
			name:    "pipeline with rulesets - push to main",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "ruleset.yml",
				Path:   "testdata",
				Type:   "",
				Branch: "main",
				Event:  "push",
			},
		},
		{
			name:    "pipeline with rulesets - tag of v1",
			failure: false,
			config: &Config{
				Action: "validate",
				File:   "ruleset.yml",
				Path:   "testdata",
				Type:   "",
				Event:  "tag",
				Tag:    "v1",
			},
		},
	}

	// run tests
	for _, test := range tests {
		isLocal := len(test.config.TemplateFiles) > 0

		err := test.config.ValidateLocal(client.WithLocal(isLocal).WithLocalTemplates(test.config.TemplateFiles))

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
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "48afb5bdc41ad69bf22588491333f7cf71135163",
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "validate",
				Org:    "github",
				Repo:   "octocat",
				Ref:    "0",
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
