// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"testing"
)

func TestConfig_ConfigFile_Empty(t *testing.T) {
	// setup tests
	tests := []struct {
		want   bool
		config *ConfigFile
	}{
		{
			want:   true,
			config: &ConfigFile{},
		},
		{
			want: false,
			config: &ConfigFile{
				API: &API{
					Address: "https://vela-server.localhost",
					Token:   "superSecretToken",
					Version: "1",
				},
				Log: &Log{
					Level: "info",
				},
				Secret: &Secret{
					Engine: "native",
					Type:   "repo",
				},
				Output: "json",
				Org:    "github",
				Repo:   "octocat",
			},
		},
		{
			want: false,
			config: &ConfigFile{
				API: &API{
					Token:   "superSecretToken",
					Version: "1",
				},
			},
		},
		{
			want: false,
			config: &ConfigFile{
				API: &API{
					Version: "1",
				},
			},
		},
		{
			want: false,
			config: &ConfigFile{
				Log: &Log{
					Level: "info",
				},
			},
		},
		{
			want: false,
			config: &ConfigFile{
				Secret: &Secret{
					Engine: "native",
					Type:   "repo",
				},
			},
		},
		{
			want: false,
			config: &ConfigFile{
				Secret: &Secret{
					Type: "repo",
				},
			},
		},
		{
			want: false,
			config: &ConfigFile{
				Output: "json",
				Org:    "github",
				Repo:   "octocat",
			},
		},
		{
			want: false,
			config: &ConfigFile{
				Org:  "github",
				Repo: "octocat",
			},
		},
		{
			want: false,
			config: &ConfigFile{
				Repo: "octocat",
			},
		},
	}

	// run tests
	for _, test := range tests {
		result := test.config.Empty()

		if result != test.want {
			t.Errorf("Empty is %t, want %t", result, test.want)
		}
	}
}
