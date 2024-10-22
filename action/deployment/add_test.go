// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/compiler/types/raw"
	"github.com/go-vela/server/mock/server"
)

func TestDeployment_Config_Add(t *testing.T) {
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
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/main",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/main",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/main",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/main",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/main",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "yaml",
				Parameters: raw.StringSliceMap{
					"foo": "bar",
					"faz": "baz",
				},
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
