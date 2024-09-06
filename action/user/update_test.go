// SPDX-License-Identifier: Apache-2.0

package user

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestRepo_Config_Update(t *testing.T) {
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
				Name:          "octocat",
				AddFavorites:  []string{"foo/bar", "foo/baz"},
				AddDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:        "",
			},
		},
		{
			failure: false,
			config: &Config{
				Name:           "octocat",
				DropFavorites:  []string{"foo/bar", "foo/baz"},
				DropDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:         "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				AddFavorites:  []string{"foo/bar", "foo/baz"},
				AddDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:        "",
			},
		},
		{
			failure: false,
			config: &Config{
				DropFavorites:  []string{"foo/bar", "foo/baz"},
				DropDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:         "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				AddFavorites:  []string{"foo/bar", "foo/baz"},
				AddDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:        "json",
			},
		},
		{
			failure: false,
			config: &Config{
				AddFavorites:  []string{"foo/bar", "foo/baz"},
				AddDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:        "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				AddFavorites:  []string{"foo/bar", "foo/baz"},
				AddDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:        "yaml",
			},
		},
		{
			failure: true,
			config: &Config{
				AddFavorites:  []string{"foo/bar", "not-found"},
				AddDashboards: []string{"c8da1302-07d6-11ea-882f-4893bca275b8"},
				Output:        "",
			},
		},
		{
			failure: true,
			config: &Config{
				AddFavorites:  []string{"foo/bar"},
				AddDashboards: []string{"0"},
				Output:        "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Update(client)

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
