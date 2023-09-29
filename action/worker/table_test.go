// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"testing"

	"github.com/go-vela/types/library"
)

func TestWorker_table(t *testing.T) {
	// setup types
	w1 := testWorker()
	w1.SetID(1)

	w2 := testWorker()
	w2.SetID(2)
	w2.SetHostname("MyWorker2")
	w2.SetAddress("myworker2.example.com")
	w2.SetRoutes([]string{"this", "that"})

	// setup tests
	tests := []struct {
		failure bool
		workers *[]library.Worker
	}{
		{
			failure: false,
			workers: &[]library.Worker{
				*w1,
				*w2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := table(test.workers)

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

func TestWorker_wideTable(t *testing.T) {
	// setup types
	w1 := testWorker()
	w1.SetID(1)
	w1.SetHostname("MyWorker")

	w2 := testWorker()
	w2.SetID(2)
	w2.SetHostname("MyWorker2")
	w2.SetAddress("myworker2.example.com")
	w2.SetRoutes([]string{"this", "that"})

	// setup tests
	tests := []struct {
		failure bool
		workers *[]library.Worker
	}{
		{
			failure: false,
			workers: &[]library.Worker{
				*w1,
				*w2,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := wideTable(test.workers)

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

// testWorker is a test helper function to create a Worker
// type with all fields set to a fake value.
func testWorker() *library.Worker {
	w := new(library.Worker)

	w.SetID(1)
	w.SetActive(true)
	w.SetBuildLimit(1)
	w.SetAddress("myworker.example.com")
	w.SetRoutes([]string{"small", "large"})
	w.SetHostname("MyWorker")
	w.SetLastCheckedIn(int64(4))

	return w
}
