// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"testing"

	api "github.com/go-vela/server/api/types"
)

func TestSecret_table(t *testing.T) {
	// setup types
	s1 := testSecret()

	s2 := testSecret()
	s2.SetID(2)
	s2.SetRepo("*")
	s2.SetType("org")

	s3 := testSecret()
	s3.SetRepo("")
	s3.SetTeam("octokitties")
	s3.SetType("shared")

	// setup tests
	tests := []struct {
		failure bool
		secrets *[]api.Secret
	}{
		{
			failure: false,
			secrets: &[]api.Secret{
				*s1,
				*s2,
				*s3,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := table(test.secrets)

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

func TestSecret_wideTable(t *testing.T) {
	// setup types
	s1 := testSecret()

	s2 := testSecret()
	s2.SetID(2)
	s2.SetRepo("*")
	s2.SetType("org")

	s3 := testSecret()
	s3.SetRepo("")
	s3.SetTeam("octokitties")
	s3.SetType("shared")
	s3.SetAllowCommand(false)
	s3.SetAllowSubstitution(false)

	// setup tests
	tests := []struct {
		failure bool
		secrets *[]api.Secret
	}{
		{
			failure: false,
			secrets: &[]api.Secret{
				*s1,
				*s2,
				*s3,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := wideTable(test.secrets)

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

// testSecret is a test helper function to create a Secret
// type with all fields set to a fake value.
func testSecret() *api.Secret {
	s := new(api.Secret)

	s.SetID(1)
	s.SetOrg("github")
	s.SetRepo("octocat")
	s.SetTeam("octokitties")
	s.SetName("foo")
	s.SetValue("bar")
	s.SetType("repo")
	s.SetImages([]string{"alpine"})
	s.GetAllowEvents().GetPush().SetBranch(true)
	s.GetAllowEvents().GetDeployment().SetCreated(false)
	s.GetAllowEvents().GetPush().SetTag(false)
	s.SetAllowCommand(true)
	s.SetAllowSubstitution(true)

	return s
}
