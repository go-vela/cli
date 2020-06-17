// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"testing"

	"github.com/go-vela/types/library"
)

func TestDeployment_table(t *testing.T) {
	// setup types
	d1 := testDeployment()

	d2 := testDeployment()
	d2.SetID(2)
	d2.SetURL("https://api.github.com/repos/github/octocat/deployments/2")

	// setup tests
	tests := []struct {
		failure bool
		steps   *[]library.Deployment
	}{
		{
			failure: false,
			steps: &[]library.Deployment{
				*d1,
				*d2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := table(test.steps)

		if test.failure {
			if err == nil {
				t.Errorf("table should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("table returned err: %v", err)
		}
	}
}

func TestDeployment_wideTable(t *testing.T) {
	// setup types
	d1 := testDeployment()

	d2 := testDeployment()
	d2.SetID(2)
	d2.SetURL("https://api.github.com/repos/github/octocat/deployments/2")

	// setup tests
	tests := []struct {
		failure bool
		steps   *[]library.Deployment
	}{
		{
			failure: false,
			steps: &[]library.Deployment{
				*d1,
				*d2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := wideTable(test.steps)

		if test.failure {
			if err == nil {
				t.Errorf("wideTable should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("wideTable returned err: %v", err)
		}
	}
}

// testDeployment is a test helper function to create a Deployment
// type with all fields set to a fake value.
func testDeployment() *library.Deployment {
	d := new(library.Deployment)

	d.SetID(1)
	d.SetRepoID(1)
	d.SetURL("https://api.github.com/repos/github/octocat/deployments/1")
	d.SetUser("octocat")
	d.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	d.SetRef("refs/heads/master")
	d.SetTask("vela-deploy")
	d.SetTarget("production")
	d.SetDescription("Deployment request from Vela")

	return d
}
