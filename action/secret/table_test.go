// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"testing"

	"github.com/go-vela/types/library"
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
		secrets *[]library.Secret
		failure bool
	}{
		{
			failure: false,
			secrets: &[]library.Secret{
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

	// setup tests
	tests := []struct {
		secrets *[]library.Secret
		failure bool
	}{
		{
			failure: false,
			secrets: &[]library.Secret{
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
func testSecret() *library.Secret {
	s := new(library.Secret)

	s.SetID(1)
	s.SetOrg("github")
	s.SetRepo("octocat")
	s.SetTeam("octokitties")
	s.SetName("foo")
	s.SetValue("bar")
	s.SetType("repo")
	s.SetImages([]string{"alpine"})
	s.SetEvents([]string{"push", "tag", "deployment"})
	s.SetAllowCommand(true)

	return s
}
