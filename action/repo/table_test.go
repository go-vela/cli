// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"testing"

	api "github.com/go-vela/server/api/types"
)

func TestRepo_table(t *testing.T) {
	// setup types
	r1 := testRepo()
	r1.GetAllowEvents().GetDeployment().SetCreated(true)
	r1.GetAllowEvents().GetPush().SetTag(true)
	r1.GetAllowEvents().GetComment().SetCreated(true)

	r2 := testRepo()
	r2.SetID(2)
	r2.SetName("octokitty")
	r2.SetFullName("github/octokitty")

	// setup tests
	tests := []struct {
		failure bool
		repos   *[]api.Repo
	}{
		{
			failure: false,
			repos: &[]api.Repo{
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
	r1.GetAllowEvents().GetDeployment().SetCreated(true)
	r1.GetAllowEvents().GetPush().SetTag(true)
	r1.GetAllowEvents().GetComment().SetCreated(true)

	r2 := testRepo()
	r2.SetID(2)
	r2.SetName("octokitty")
	r2.SetFullName("github/octokitty")

	// setup tests
	tests := []struct {
		failure bool
		repos   *[]api.Repo
	}{
		{
			failure: false,
			repos: &[]api.Repo{
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
func testRepo() *api.Repo {
	r := new(api.Repo)

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
	r.GetAllowEvents().GetPullRequest().SetOpened(true)
	r.GetAllowEvents().GetPullRequest().SetSynchronize(true)
	r.GetAllowEvents().GetPullRequest().SetEdited(true)
	r.GetAllowEvents().GetPush().SetBranch(true)
	r.GetAllowEvents().GetDeployment().SetCreated(false)
	r.GetAllowEvents().GetPush().SetTag(false)
	r.GetAllowEvents().GetComment().SetCreated(false)

	return r
}
