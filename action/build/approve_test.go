// SPDX-License-Identifier: Apache-2.0

package build

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestBuild_Config_Approve(t *testing.T) {
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
				Action: "approve",
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "approve",
				Org:    "github",
				Repo:   "octocat",
				Number: 0,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Approve(client)

		if test.failure {
			if err == nil {
				t.Errorf("Approve should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Approve returned err: %v", err)
		}
	}
}
