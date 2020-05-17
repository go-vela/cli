// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/go-vela/sdk-go/vela"
)

func TestBuild_Build_Get(t *testing.T) {
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
		build   *Build
	}{
		{
			failure: false,
			build: &Build{
				Action:  getAction,
				Org:     "github",
				Repo:    "octocat",
				Number:  1,
				Page:    1,
				PerPage: 10,
				Output:  "json",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.build.Get(client)

		if test.failure {
			if err == nil {
				t.Errorf("Get should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Get returned err: %v", err)
		}
	}
}

func TestBuild_Build_Restart(t *testing.T) {
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
		build   *Build
	}{
		{
			failure: false,
			build: &Build{
				Action: restartAction,
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.build.Restart(client)

		if test.failure {
			if err == nil {
				t.Errorf("Restart should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Restart returned err: %v", err)
		}
	}
}

func TestBuild_Build_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		build   *Build
	}{
		{
			failure: false,
			build: &Build{
				Action:  getAction,
				Org:     "github",
				Repo:    "octocat",
				Page:    1,
				PerPage: 10,
				Output:  "json",
			},
		},
		{
			failure: false,
			build: &Build{
				Action: restartAction,
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: false,
			build: &Build{
				Action: viewAction,
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: true,
			build: &Build{
				Action: viewAction,
				Org:    "",
				Repo:   "octocat",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: true,
			build: &Build{
				Action: viewAction,
				Org:    "github",
				Repo:   "",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: true,
			build: &Build{
				Action: viewAction,
				Org:    "github",
				Repo:   "octocat",
				Number: 0,
				Output: "json",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.build.Validate()

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

func TestBuild_Build_View(t *testing.T) {
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
		build   *Build
	}{
		{
			failure: false,
			build: &Build{
				Action: viewAction,
				Org:    "github",
				Repo:   "octocat",
				Number: 1,
				Output: "json",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.build.View(client)

		if test.failure {
			if err == nil {
				t.Errorf("View should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("View returned err: %v", err)
		}
	}
}
