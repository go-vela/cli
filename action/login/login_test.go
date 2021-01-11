// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestLogin_Config_Login(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "Vela CLI", nil)
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
				Action: "login",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Login(client, "http://localhost:8080")

		if test.failure {
			if err == nil {
				t.Errorf("Login should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Login returned err: %v", err)
		}
	}
}
