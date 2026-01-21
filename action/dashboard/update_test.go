// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestDashboard_Config_Update(t *testing.T) {
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
				Action: "update",
				ID:     "c8da1302-07d6-11ea-882f-4893bca275b8",
				Name:   "my-dashboard",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:    "update",
				ID:        "c8da1302-07d6-11ea-882f-4893bca275b8",
				AddRepos:  []string{"github/octocat"},
				AddAdmins: []string{"octocat"},
			},
		},
		{
			failure: false,
			config: &Config{
				Action:     "update",
				ID:         "c8da1302-07d6-11ea-882f-4893bca275b8",
				DropRepos:  []string{"github/octocat"},
				DropAdmins: []string{"octocat"},
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "update",
				ID:          "c8da1302-07d6-11ea-882f-4893bca275b8",
				Branches:    []string{"main", "dev"},
				Events:      []string{"push", "tag"},
				TargetRepos: []string{"github/octocat"},
				Output:      "json",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Update(t.Context(), client)

		if test.failure {
			if err == nil {
				t.Errorf("Update should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Update returned err: %v", err)
		}
	}
}
