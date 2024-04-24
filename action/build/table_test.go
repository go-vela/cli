// SPDX-License-Identifier: Apache-2.0

package build

import (
	"testing"

	api "github.com/go-vela/server/api/types"
)

func TestBuild_table(t *testing.T) {
	// setup types
	b1 := testBuild()
	b1.SetStatus("success")
	b1.SetStarted(1563474214)
	b1.SetFinished(1563474224)

	b2 := testBuild()
	b2.SetID(2)
	b2.SetNumber(2)
	b2.SetMessage("Second commit...")

	// setup tests
	tests := []struct {
		failure bool
		builds  *[]api.Build
	}{
		{
			failure: false,
			builds: &[]api.Build{
				*b1,
				*b2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := table(test.builds)

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

func TestBuild_wideTable(t *testing.T) {
	// setup types
	b1 := testBuild()
	b1.SetStatus("success")
	b1.SetStarted(1563474214)
	b1.SetFinished(1563474224)

	b2 := testBuild()
	b2.SetID(2)
	b2.SetNumber(2)
	b2.SetMessage("Second commit...")

	// setup tests
	tests := []struct {
		failure bool
		builds  *[]api.Build
	}{
		{
			failure: false,
			builds: &[]api.Build{
				*b1,
				*b2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := wideTable(test.builds)

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

// testBuild is a test helper function to create a Build
// type with all fields set to a fake value.
func testBuild() *api.Build {
	r := new(api.Repo)
	r.SetID(1)

	b := new(api.Build)

	b.SetID(1)
	b.SetRepo(r)
	b.SetNumber(1)
	b.SetParent(1)
	b.SetEvent("push")
	b.SetStatus("running")
	b.SetError("")
	b.SetEnqueued(1563474077)
	b.SetCreated(1563474076)
	b.SetStarted(1563474078)
	b.SetFinished(1563474079)
	b.SetDeploy("")
	b.SetClone("https://github.com/github/octocat.git")
	b.SetSource("https://github.com/github/octocat/48afb5bdc41ad69bf22588491333f7cf71135163")
	b.SetTitle("push received from https://github.com/github/octocat")
	b.SetMessage("First commit...")
	b.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	b.SetSender("OctoKitty")
	b.SetAuthor("OctoKitty")
	b.SetEmail("OctoKitty@github.com")
	b.SetLink("https://example.company.com/github/octocat/1")
	b.SetBranch("main")
	b.SetRef("refs/heads/main")
	b.SetBaseRef("")
	b.SetHost("example.company.com")
	b.SetRuntime("docker")
	b.SetDistribution("linux")

	return b
}
