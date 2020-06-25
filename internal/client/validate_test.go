// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package client

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"
)

func TestClient_validate(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup tests
	tests := []struct {
		failure bool
		address string
		token   string
	}{
		{
			failure: false,
			address: s.URL,
			token:   "superSecretToken",
		},
		{
			failure: true,
			address: "",
			token:   "superSecretToken",
		},
		{
			failure: true,
			address: s.URL,
			token:   "",
		},
	}

	// run tests
	for _, test := range tests {
		err := validate(test.address, test.token)

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
