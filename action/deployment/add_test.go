// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-vela/mock/server"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/raw"
)

func TestDeployment_Config_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, nil)
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
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/master",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/master",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/master",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/master",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/master",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "yaml",
				Payload:     []string{"foo=test1", "bar=test2"},
			},
		},
		{
			failure: true,
			config: &Config{
				Action:      "add",
				Org:         "github",
				Repo:        "octocat",
				Description: "Deployment request from Vela",
				Ref:         "refs/heads/master",
				Target:      "production",
				Task:        "deploy:vela",
				Output:      "yaml",
				Payload:     []string{"badinput"},
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

func Test_parseKeyValue(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    raw.StringSliceMap
		wantErr bool
	}{
		{"valid input", args{input: []string{"foo=test1", "bar=test2"}}, raw.StringSliceMap{"foo": "test1", "bar": "test2"}, false},
		{"invalid input", args{input: []string{"foo=test1", "badinput"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseKeyValue(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseKeyValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKeyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
