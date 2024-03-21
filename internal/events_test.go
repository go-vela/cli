// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"testing"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/library/actions"
	"github.com/google/go-cmp/cmp"
)

func Test_PopulateEvents(t *testing.T) {
	// setup types
	tBool := true

	// setup tests
	tests := []struct {
		name   string
		events []string
		want   *library.Events
	}{
		{
			name:   "general events",
			events: []string{"push", "pull_request", "tag", "deploy", "comment", "delete", "schedule"},
			want: &library.Events{
				Push: &actions.Push{
					Branch:       &tBool,
					Tag:          &tBool,
					DeleteBranch: &tBool,
					DeleteTag:    &tBool,
				},
				PullRequest: &actions.Pull{
					Opened:      &tBool,
					Reopened:    &tBool,
					Synchronize: &tBool,
				},
				Deployment: &actions.Deploy{
					Created: &tBool,
				},
				Comment: &actions.Comment{
					Created: &tBool,
					Edited:  &tBool,
				},
				Schedule: &actions.Schedule{
					Run: &tBool,
				},
			},
		},
		{
			name:   "action specific events",
			events: []string{"push:branch", "push:tag", "pull_request:opened", "pull_request:edited", "deployment:created", "comment:created", "delete:branch", "delete:tag", "schedule:run"},
			want: &library.Events{
				Push: &actions.Push{
					Branch:       &tBool,
					Tag:          &tBool,
					DeleteBranch: &tBool,
					DeleteTag:    &tBool,
				},
				PullRequest: &actions.Pull{
					Opened: &tBool,
					Edited: &tBool,
				},
				Deployment: &actions.Deploy{
					Created: &tBool,
				},
				Comment: &actions.Comment{
					Created: &tBool,
				},
				Schedule: &actions.Schedule{
					Run: &tBool,
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		got := PopulateEvents(test.events)

		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("PopulateEvents failed for %s mismatch (-want +got):\n%s", test.name, diff)
		}
	}
}
