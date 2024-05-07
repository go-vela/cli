// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestDashboard_Config_Add(t *testing.T) {
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
				Action:    "add",
				Name:      "my-dashboard",
				AddRepos:  []string{"github/octocat"},
				AddAdmins: []string{"octocat"},
				Branches:  []string{"main", "dev"},
				Events:    []string{"push", "tag"},
				Output:    "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "add",
				Name:   "my-dashboard",
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
