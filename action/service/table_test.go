// SPDX-License-Identifier: Apache-2.0

package service

import (
	"testing"

	api "github.com/go-vela/server/api/types"
)

func TestService_table(t *testing.T) {
	// setup types
	s1 := testService()
	s1.SetStatus("success")
	s1.SetStarted(1563474214)
	s1.SetFinished(1563474224)

	s2 := testService()
	s2.SetID(2)
	s2.SetNumber(2)

	// setup tests
	tests := []struct {
		failure bool
		steps   *[]api.Service
	}{
		{
			failure: false,
			steps: &[]api.Service{
				*s1,
				*s2,
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

func TestService_wideTable(t *testing.T) {
	// setup types
	s1 := testService()
	s1.SetStatus("success")
	s1.SetStarted(1563474214)
	s1.SetFinished(1563474224)

	s2 := testService()
	s2.SetID(2)
	s2.SetNumber(2)

	// setup tests
	tests := []struct {
		failure bool
		steps   *[]api.Service
	}{
		{
			failure: false,
			steps: &[]api.Service{
				*s1,
				*s2,
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

// testService is a test helper function to create a Service
// type with all fields set to a fake value.
func testService() *api.Service {
	s := new(api.Service)

	s.SetID(1)
	s.SetBuildID(1)
	s.SetRepoID(1)
	s.SetNumber(1)
	s.SetName("postgres")
	s.SetImage("postgres:12-alpine")
	s.SetStatus("running")
	s.SetExitCode(0)
	s.SetCreated(1563474076)
	s.SetStarted(1563474078)
	s.SetFinished(1563474079)
	s.SetHost("example.company.com")
	s.SetRuntime("docker")
	s.SetDistribution("linux")

	return s
}
