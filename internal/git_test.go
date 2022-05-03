// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package internal

import (
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func TestGit_GetGitConfigOrg(t *testing.T) {
	testSetup(t)
	// setup tests
	tests := []struct {
		expect string
		input  string
	}{
		{
			expect: "go-vela",
			input:  "./testdata/project1/",
		},
		{
			expect: "OctoKitty",
			input:  "./testdata/project2/",
		},
		{
			expect: "testing",
			input:  "./testdata/project3/",
		},
		{
			expect: "http-test",
			input:  "./testdata/project4/",
		},
	}

	// run tests
	for _, test := range tests {
		got := GetGitConfigOrg(test.input)
		if got != test.expect {
			t.Errorf("GetGitConfigOrg returned org: %s\n Expected org: %s\n", got, test.expect)
		}
	}

	testTearDown()
}

func TestGit_GetGitConfigRepo(t *testing.T) {
	testSetup(t)
	// setup tests
	tests := []struct {
		expect string
		input  string
	}{
		{
			expect: "api-handler",
			input:  "./testdata/project1/",
		},
		{ // map
			expect: "catCastle",
			input:  "./testdata/project2/",
		},
		{
			expect: "test-test",
			input:  "./testdata/project3/",
		},
		{
			expect: "test-repo",
			input:  "./testdata/project4/",
		},
	}

	// run tests
	for _, test := range tests {
		got := GetGitConfigRepo(test.input)
		if got != test.expect {
			t.Errorf("GetGitConfigRepo returned repo: %s\n Expected repo: %s\n", got, test.expect)
		}
	}

	testTearDown()
}

func testSetup(t *testing.T) {
	// setup configs
	r1, err := git.PlainInit("./testdata/project1/", false)
	if err != nil {
		t.Errorf("Failed to init repo for project 1")
	}

	_, err = r1.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{"git@github.com:go-vela/api-handler.git"},
	})
	if err != nil {
		t.Errorf("Could not create remote")
	}

	r2, err := git.PlainInit("./testdata/project2/", false)
	if err != nil {
		t.Errorf("Failed to init repo for project 2")
	}

	_, err = r2.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{"git@github.com:OctoKitty/catCastle"},
	})
	if err != nil {
		t.Errorf("Could not create remote")
	}

	r3, err := git.PlainInit("./testdata/project3/", false)
	if err != nil {
		t.Errorf("Failed to init repo for project 3")
	}

	_, err = r3.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{"ssh://git@github.com/testing/test-test.git"},
	})
	if err != nil {
		t.Errorf("Could not create remote")
	}

	r4, err := git.PlainInit("./testdata/project4/", false)
	if err != nil {
		t.Errorf("Failed to init repo for project 4")
	}

	_, err = r4.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{"https://github.com/http-test/test-repo"},
	})
	if err != nil {
		t.Errorf("Could not create remote")
	}
}

func testTearDown() {
	os.RemoveAll("./testdata/project1/.git")
	os.RemoveAll("./testdata/project2/.git")
	os.RemoveAll("./testdata/project3/.git")
	os.RemoveAll("./testdata/project4/.git")
}
