// SPDX-License-Identifier: Apache-2.0

package hook

import (
	"testing"
	"time"

	api "github.com/go-vela/server/api/types"
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
		steps   *[]api.Hook
	}{
		{
			failure: false,
			steps: &[]api.Hook{
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
		steps   *[]api.Hook
	}{
		{
			failure: false,
			steps: &[]api.Hook{
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
func testHook() *api.Hook {
	h := new(api.Hook)

	h.SetID(1)
	h.SetNumber(1)
	h.SetSourceID("c8da1302-07d6-11ea-882f-4893bca275b8")
	h.SetCreated(time.Now().UTC().Unix())
	h.SetHost("github.com")
	h.SetEvent("push")
	h.SetBranch("main")
	h.SetError("")
	h.SetStatus("success")
	h.SetLink("https://github.com/github/octocat/settings/hooks/1")

	return h
}
