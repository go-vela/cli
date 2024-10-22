// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"testing"

	api "github.com/go-vela/server/api/types"
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
		steps   *[]api.Deployment
	}{
		{
			failure: false,
			steps: &[]api.Deployment{
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
		steps   *[]api.Deployment
	}{
		{
			failure: false,
			steps: &[]api.Deployment{
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
func testDeployment() *api.Deployment {
	d := new(api.Deployment)

	d.SetID(1)
	d.SetURL("https://api.github.com/repos/github/octocat/deployments/1")
	d.SetCreatedBy("octocat")
	d.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	d.SetRef("refs/heads/main")
	d.SetTask("vela-deploy")
	d.SetTarget("production")
	d.SetDescription("Deployment request from Vela")

	return d
}
