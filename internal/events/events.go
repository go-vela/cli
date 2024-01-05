// SPDX-License-Identifier: Apache-2.0

package events

import (
	"fmt"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/library/actions"
)

// Alternate constants for webhook events.
const (
	AlternateDeploy = "deploy"
	AlternatePull   = "pull"
)

// Populate is a helper function designed to fill both the legacy `Allow<_>` fields
// as well as the `AllowEvents` struct with the correct values based on a slice input.
func Populate(events []string) (*library.Events, error) {
	result := new(library.Events)
	push := new(actions.Push)
	pull := new(actions.Pull)
	comment := new(actions.Comment)
	deploy := new(actions.Deploy)

	// iterate through all events provided
	for _, event := range events {
		switch event {
		case constants.EventPush, constants.EventPush + ":branch":
			push.SetBranch(true)
		case constants.EventTag, constants.EventPush + ":" + constants.EventTag:
			push.SetTag(true)
		case constants.EventPull, AlternatePull:
			pull.SetOpened(true)
			pull.SetReopened(true)
			pull.SetSynchronize(true)
		case constants.EventDeploy, AlternateDeploy, constants.EventDeploy + ":" + constants.ActionCreated:
			deploy.SetCreated(true)
		case constants.EventComment:
			comment.SetCreated(true)
			comment.SetEdited(true)
		case constants.EventPull + ":" + constants.ActionOpened:
			pull.SetOpened(true)
		case constants.EventPull + ":" + constants.ActionEdited:
			pull.SetEdited(true)
		case constants.EventPull + ":" + constants.ActionSynchronize:
			pull.SetSynchronize(true)
		case constants.EventPull + ":" + constants.ActionReopened:
			pull.SetReopened(true)
		case constants.EventComment + ":" + constants.ActionCreated:
			comment.SetCreated(true)
		case constants.EventComment + ":" + constants.ActionEdited:
			comment.SetEdited(true)
		default:
			return nil, fmt.Errorf("invalid event provided: %s", event)
		}
	}

	result.SetPush(push)
	result.SetPullRequest(pull)
	result.SetDeployment(deploy)
	result.SetComment(comment)

	return result, nil
}
