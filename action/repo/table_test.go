// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"testing"

	"github.com/go-vela/types/library"
)

func TestRepo_table(t *testing.T) {
	// setup types
	r1 := testRepo()
	r1.SetAllowDeploy(true)
	r1.SetAllowTag(true)
	r1.SetAllowComment(true)

	r2 := testRepo()
	r2.SetID(2)
	r2.SetName("octokitty")
	r2.SetFullName("github/octokitty")

	// setup tests
	tests := []struct {
		failure bool
		repos   *[]library.Repo
	}{
		{
			failure: false,
			repos: &[]library.Repo{
				*r1,
				*r2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := table(test.repos)

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

func TestRepo_wideTable(t *testing.T) {
	// setup types
	r1 := testRepo()
	r1.SetAllowDeploy(true)
	r1.SetAllowTag(true)
	r1.SetAllowComment(true)

	r2 := testRepo()
	r2.SetID(2)
	r2.SetName("octokitty")
	r2.SetFullName("github/octokitty")

	// setup tests
	tests := []struct {
		failure bool
		repos   *[]library.Repo
	}{
		{
			failure: false,
			repos: &[]library.Repo{
				*r1,
				*r2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := wideTable(test.repos)

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

// testRepo is a test helper function to create a Repo
// type with all fields set to a fake value.
func testRepo() *library.Repo {
	r := new(library.Repo)

	r.SetID(1)
	r.SetOrg("github")
	r.SetName("octocat")
	r.SetFullName("github/octocat")
	r.SetLink("https://github.com/github/octocat")
	r.SetClone("https://github.com/github/octocat.git")
	r.SetBranch("main")
	r.SetTimeout(30)
	r.SetVisibility("public")
	r.SetPrivate(false)
	r.SetTrusted(false)
	r.SetActive(true)
	r.SetAllowPull(true)
	r.SetAllowPush(true)
	r.SetAllowDeploy(false)
	r.SetAllowTag(false)
	r.SetAllowComment(false)

	return r
}
