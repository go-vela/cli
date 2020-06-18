// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

import (
	"testing"
	"time"

	"github.com/go-vela/types/library"
)

func TestHook_table(t *testing.T) {
	// setup types
	h1 := testHook()

	h2 := testHook()
	h2.SetID(2)
	h2.SetNumber(2)

	// setup tests
	tests := []struct {
		failure bool
		steps   *[]library.Hook
	}{
		{
			failure: false,
			steps: &[]library.Hook{
				*h1,
				*h2,
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

func TestHook_wideTable(t *testing.T) {
	// setup types
	h1 := testHook()

	h2 := testHook()
	h2.SetID(2)
	h2.SetNumber(2)

	// setup tests
	tests := []struct {
		failure bool
		steps   *[]library.Hook
	}{
		{
			failure: false,
			steps: &[]library.Hook{
				*h1,
				*h2,
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

// testHook is a test helper function to create a Hook
// type with all fields set to a fake value.
func testHook() *library.Hook {
	h := new(library.Hook)

	h.SetID(1)
	h.SetRepoID(1)
	h.SetBuildID(1)
	h.SetNumber(1)
	h.SetSourceID("c8da1302-07d6-11ea-882f-4893bca275b8")
	h.SetCreated(time.Now().UTC().Unix())
	h.SetHost("github.com")
	h.SetEvent("push")
	h.SetBranch("master")
	h.SetError("")
	h.SetStatus("success")
	h.SetLink("https://github.com/github/octocat/settings/hooks/1")

	return h
}
