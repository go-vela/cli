// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/library/actions"
	"github.com/google/go-cmp/cmp"

	"github.com/go-vela/sdk-go/vela"
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
				Output:       "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Add(client)

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

func TestRepo_populateEvents(t *testing.T) {
	// setup types
	tBool := true
	fBool := false

	// setup tests
	tests := []struct {
		name   string
		events []string
		want   *library.Repo
	}{
		{
			name:   "happy path legacy events",
			events: []string{"push", "pull_request", "tag", "deploy", "comment"},
			want: &library.Repo{
				AllowPush:    &tBool,
				AllowPull:    &tBool,
				AllowTag:     &tBool,
				AllowDeploy:  &tBool,
				AllowComment: &tBool,
				AllowEvents: &library.Events{
					Push: &actions.Push{
						Branch: &tBool,
						Tag:    &tBool,
					},
					PullRequest: &actions.Pull{
						Opened:      &tBool,
						Reopened:    &tBool,
						Synchronize: &tBool,
					},
					Deployment: &actions.Deploy{
						Created: &tBool,
					},
					Comment: &actions.Comment{
						Created: &tBool,
						Edited:  &tBool,
					},
				},
			},
		},
		{
			name:   "action specific",
			events: []string{"push:branch", "push:tag", "pull_request:opened", "pull_request:edited", "deployment:created", "comment:created"},
			want: &library.Repo{
				AllowPush:    &tBool,
				AllowPull:    &fBool,
				AllowTag:     &tBool,
				AllowDeploy:  &tBool,
				AllowComment: &fBool,
				AllowEvents: &library.Events{
					Push: &actions.Push{
						Branch: &tBool,
						Tag:    &tBool,
					},
					PullRequest: &actions.Pull{
						Opened: &tBool,
						Edited: &tBool,
					},
					Deployment: &actions.Deploy{
						Created: &tBool,
					},
					Comment: &actions.Comment{
						Created: &tBool,
					},
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		repo := new(library.Repo)
		populateEvents(repo, test.events)

		if diff := cmp.Diff(test.want, repo); diff != "" {
			t.Errorf("populateEvents failed for %s mismatch (-want +got):\n%s", test.name, diff)
		}
	}
}
