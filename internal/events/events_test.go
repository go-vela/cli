// SPDX-License-Identifier: Apache-2.0

package events

import (
	"testing"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/library/actions"
	"github.com/google/go-cmp/cmp"
)

func TestEvents_Populate(t *testing.T) {
	// setup types
	tBool := true

	// setup tests
	tests := []struct {
		name    string
		events  []string
		want    *library.Events
		wantErr bool
	}{
		{
			name:   "happy path legacy events",
			events: []string{"push", "pull_request", "tag", "deploy", "comment"},
			want: &library.Events{
				Push: &actions.Push{
					Branch: &tBool,
					Tag:    &tBool,
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
			},
		},
		{
			name:   "action specific",
			events: []string{"push:branch", "push:tag", "pull_request:opened", "pull_request:edited", "deployment:created", "comment:created"},
			want: &library.Events{
				Push: &actions.Push{
					Branch: &tBool,
					Tag:    &tBool,
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
			},
		},
		{
			name:    "invalid event",
			events:  []string{"pull_request:edited", "invalid", "comment:created"},
			wantErr: true,
		},
	}

	// run tests
	for _, test := range tests {
		got, err := Populate(test.events)
		if err != nil && !test.wantErr {
			t.Errorf("Populate returned err: %s", err)
		}

		if err == nil && test.wantErr {
			t.Errorf("Populate should have returned error")
		}

		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("populateEvents failed for %s mismatch (-want +got):\n%s", test.name, diff)
		}
	}
}
