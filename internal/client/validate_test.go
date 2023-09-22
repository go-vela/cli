// SPDX-License-Identifier: Apache-2.0

package client

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/server/mock/server"
)

func TestClient_validate(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup tests
	tests := []struct {
		failure      bool
		address      string
		token        string
		accessToken  string
		refreshToken string
	}{
		{
			failure:      false,
			address:      s.URL,
			token:        "superSecretToken",
			accessToken:  "",
			refreshToken: "",
		},
		{
			failure:      true,
			address:      "",
			token:        "",
			accessToken:  "superSecretAccessToken",
			refreshToken: "superSecretRefreshToken",
		},
		{
			failure:      true,
			address:      "",
			token:        "superSecretToken",
			accessToken:  "",
			refreshToken: "",
		},
		{
			failure:      true,
			address:      s.URL,
			token:        "",
			accessToken:  "",
			refreshToken: "",
		},
	}

	// run tests
	for _, test := range tests {
		err := validate(test.address, test.token, test.accessToken, test.refreshToken)

		if test.failure {
			if err == nil {
				t.Errorf("validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("validate returned err: %v", err)
		}
	}
}
