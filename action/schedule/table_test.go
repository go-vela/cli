// SPDX-License-Identifier: Apache-2.0

package schedule

import (
	"testing"
	"time"

	api "github.com/go-vela/server/api/types"
)

func TestSchedule_table(t *testing.T) {
	// setup types
	_scheduleOne := testSchedule()

	_scheduleTwo := testSchedule()
	_scheduleTwo.SetID(2)
	_scheduleTwo.SetName("bar")

	// setup tests
	tests := []struct {
		name      string
		failure   bool
		schedules *[]api.Schedule
	}{
		{
			name:      "success",
			failure:   false,
			schedules: &[]api.Schedule{*_scheduleOne, *_scheduleTwo},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := table(test.schedules)

			if test.failure {
				if err == nil {
					t.Errorf("table for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("table for %s returned err: %v", test.name, err)
			}
		})
	}
}

func TestSchedule_wideTable(t *testing.T) {
	// setup types
	_scheduleOne := testSchedule()

	_scheduleTwo := testSchedule()
	_scheduleTwo.SetID(2)
	_scheduleTwo.SetName("bar")

	// setup tests
	tests := []struct {
		name      string
		failure   bool
		schedules *[]api.Schedule
	}{
		{
			name:      "success",
			failure:   false,
			schedules: &[]api.Schedule{*_scheduleOne, *_scheduleTwo},
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := wideTable(test.schedules)

			if test.failure {
				if err == nil {
					t.Errorf("wideTable for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("wideTable for %s returned err: %v", test.name, err)
			}
		})
	}
}

// testSchedule is a test helper function to create a Schedule type with all fields set to a fake value.
func testSchedule() *api.Schedule {
	s := new(api.Schedule)
	s.SetID(1)
	s.SetActive(true)
	s.SetName("nightly")
	s.SetEntry("0 0 * * *")
	s.SetCreatedAt(time.Now().UTC().Unix())
	s.SetCreatedBy("user1")
	s.SetUpdatedAt(time.Now().Add(time.Hour * 1).UTC().Unix())
	s.SetUpdatedBy("user2")
	s.SetScheduledAt(time.Now().Add(time.Hour * 2).UTC().Unix())
	s.SetBranch("main")

	return s
}
