// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"testing"

	"github.com/go-vela/types/library"
)

func TestPipeline_table(t *testing.T) {
	// setup types
	p1 := testPipeline()

	p2 := testPipeline()
	p2.SetID(2)
	p2.SetCommit("a49aaf4afae6431a79239c95247a2b169fd9f067")

	// setup tests
	tests := []struct {
		name      string
		failure   bool
		pipelines *[]library.Pipeline
	}{
		{
			name:    "success",
			failure: false,
			pipelines: &[]library.Pipeline{
				*p1,
				*p2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := table(test.pipelines)

			if test.failure {
				if err == nil {
					t.Errorf("table should have returned err")
				}

				return
			}

			if err != nil {
				t.Errorf("table returned err: %v", err)
			}
		})
	}
}

func TestPipeline_wideTable(t *testing.T) {
	// setup types
	p1 := testPipeline()

	p2 := testPipeline()
	p2.SetID(2)
	p2.SetCommit("a49aaf4afae6431a79239c95247a2b169fd9f067")

	// setup tests
	tests := []struct {
		name      string
		failure   bool
		pipelines *[]library.Pipeline
	}{
		{
			name:    "success",
			failure: false,
			pipelines: &[]library.Pipeline{
				*p1,
				*p2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := wideTable(test.pipelines)

			if test.failure {
				if err == nil {
					t.Errorf("wideTable should have returned err")
				}

				return
			}

			if err != nil {
				t.Errorf("wideTable returned err: %v", err)
			}
		})
	}
}

// testPipeline is a test helper function to create a Pipeline
// type with all fields set to a fake value.
func testPipeline() *library.Pipeline {
	p := new(library.Pipeline)

	p.SetID(1)
	p.SetRepoID(1)
	p.SetCommit("48afb5bdc41ad69bf22588491333f7cf71135163")
	p.SetFlavor("large")
	p.SetPlatform("docker")
	p.SetRef("refs/heads/master")
	p.SetRef("yaml")
	p.SetVersion("1")
	p.SetExternalSecrets(false)
	p.SetInternalSecrets(false)
	p.SetServices(true)
	p.SetStages(false)
	p.SetSteps(true)
	p.SetTemplates(false)

	return p
}
