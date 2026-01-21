// SPDX-License-Identifier: Apache-2.0

package build

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestBuild_Config_Cancel(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "", nil)
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
				Action: "cancel",
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "cancel",
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "cancel",
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "cancel",
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "cancel",
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "cancel",
				Org:    "github",
				Repo:   "octocat",
				Number: 0,
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Cancel(t.Context(), client)

		if test.failure {
			if err == nil {
				t.Errorf("Cancel should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Cancel returned err: %v", err)
		}
	}
}
