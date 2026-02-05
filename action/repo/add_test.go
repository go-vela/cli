// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestRepo_Config_Add(t *testing.T) {
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
				Action:       "add",
				Org:          "github",
				Name:         "octocat",
				Link:         "https://github.com/github/octocat",
				Clone:        "https://github.com/github/octocat.git",
				Branch:       "main",
				BuildLimit:   10,
				Timeout:      60,
				Counter:      0,
				Visibility:   "public",
				Private:      false,
				Trusted:      false,
				Active:       true,
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
				PipelineType: "yaml",
				ApproveBuild: "fork-always",
				Output:       "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:       "add",
				Org:          "github",
				Name:         "octocat",
				Link:         "https://github.com/github/octocat",
				Clone:        "https://github.com/github/octocat.git",
				Branch:       "main",
				BuildLimit:   10,
				Timeout:      60,
				Counter:      0,
				Visibility:   "public",
				Private:      false,
				Trusted:      false,
				Active:       true,
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
				PipelineType: "yaml",
				Output:       "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:       "add",
				Org:          "github",
				Name:         "octocat",
				Link:         "https://github.com/github/octocat",
				Clone:        "https://github.com/github/octocat.git",
				Branch:       "main",
				BuildLimit:   10,
				Timeout:      60,
				Visibility:   "public",
				Private:      false,
				Trusted:      false,
				Active:       true,
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
				PipelineType: "yaml",
				ApproveBuild: "fork-no-write",
				Output:       "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:       "add",
				Org:          "github",
				Name:         "octocat",
				Link:         "https://github.com/github/octocat",
				Clone:        "https://github.com/github/octocat.git",
				Branch:       "main",
				BuildLimit:   10,
				Timeout:      60,
				Counter:      0,
				Visibility:   "public",
				Private:      false,
				Trusted:      false,
				Active:       true,
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
				PipelineType: "yaml",
				Output:       "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:       "add",
				Org:          "github",
				Name:         "octocat",
				Link:         "https://github.com/github/octocat",
				Clone:        "https://github.com/github/octocat.git",
				Branch:       "main",
				BuildLimit:   10,
				Timeout:      60,
				Counter:      0,
				Visibility:   "public",
				Private:      false,
				Trusted:      false,
				Active:       true,
				Events:       []string{"push", "pull_request", "comment", "deployment", "tag"},
				PipelineType: "yaml",
				ApproveBuild: "never",
				Output:       "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Add(t.Context(), client)

		if test.failure {
			if err == nil {
				t.Errorf("Add should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Add returned err: %v", err)
		}
	}
}
