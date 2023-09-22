// SPDX-License-Identifier: Apache-2.0

package schedule

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSchedule_Config_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create test client: %v", err)
	}

	// setup tests
	tests := []struct {
		name    string
		failure bool
		config  *Config
	}{
		{
			name:    "success with empty output",
			failure: false,
			config: &Config{
				Action: "add",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "",
				Branch: "main",
			},
		},
		{
			name:    "success with dump output",
			failure: false,
			config: &Config{
				Action: "add",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "dump",
				Branch: "main",
			},
		},
		{
			name:    "success with json output",
			failure: false,
			config: &Config{
				Action: "add",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "json",
				Branch: "main",
			},
		},
		{
			name:    "success with spew output",
			failure: false,
			config: &Config{
				Action: "add",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "spew",
				Branch: "main",
			},
		},
		{
			name:    "success with yaml output",
			failure: false,
			config: &Config{
				Action: "add",
				Org:    "github",
				Repo:   "octocat",
				Active: true,
				Name:   "foo",
				Entry:  "@weekly",
				Output: "yaml",
				Branch: "main",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err = test.config.Add(client)

			if test.failure {
				if err == nil {
					t.Errorf("Add for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Add for %s returned err: %v", test.name, err)
			}
		})
	}
}
