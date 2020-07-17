// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

import (
	"testing"
)

func TestLogin_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:   "login",
				Username: "octocat",
				Password: "superSecretPassword",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:   "login",
				Username: "octocat",
				Password: "superSecretPassword",
				OTP:      "123456",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "login",
				Password: "superSecretPassword",
				OTP:      "123456",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "login",
				Username: "octocat",
				OTP:      "123456",
			},
		},
		{
			failure: true,
			config: &Config{
				Action:   "login",
				Username: "octocat",
				Password: "superSecretPassword",
				Retry:    true,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}
