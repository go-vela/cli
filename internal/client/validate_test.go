// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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
		address      string
		token        string
		accessToken  string
		refreshToken string
		failure      bool
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
