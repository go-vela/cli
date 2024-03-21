// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/library/actions"
)

const (
	// AlternatePull is an alternate name for the pull_request event.
	AlternatePull = "pull"
	// AlternateDeploy is an alternate name for the deployment event.
	AlternateDeploy = "deploy"
)

// PopulateEvents is a helper function designed to fill both the legacy `Allow<_>` fields
// as well as the `AllowEvents` struct with the correct values based on a slice input.
func PopulateEvents(events []string) *library.Events {
	result := new(library.Events)
	push := new(actions.Push)
	pull := new(actions.Pull)
	comment := new(actions.Comment)
	deploy := new(actions.Deploy)
	schedule := new(actions.Schedule)

	// iterate through all events provided
	for _, event := range events {
		switch event {
		// push actions
		case constants.EventPush, constants.EventPush + ":branch":
			push.SetBranch(true)
		case constants.EventTag, constants.EventPush + ":" + constants.EventTag:
			push.SetTag(true)
		case constants.EventDelete + ":" + constants.ActionBranch:
			push.SetDeleteBranch(true)
		case constants.EventDelete + ":" + constants.ActionTag:
			push.SetDeleteTag(true)
		case constants.EventDelete:
			push.SetDeleteBranch(true)
			push.SetDeleteTag(true)

		// pull_request actions
		case constants.EventPull, AlternatePull:
			pull.SetOpened(true)
			pull.SetReopened(true)
			pull.SetSynchronize(true)
		case constants.EventPull + ":" + constants.ActionOpened:
			pull.SetOpened(true)
		case constants.EventPull + ":" + constants.ActionEdited:
			pull.SetEdited(true)
		case constants.EventPull + ":" + constants.ActionSynchronize:
			pull.SetSynchronize(true)
		case constants.EventPull + ":" + constants.ActionReopened:
			pull.SetReopened(true)

		// deployment actions
		case constants.EventDeploy, AlternateDeploy, constants.EventDeploy + ":" + constants.ActionCreated:
			deploy.SetCreated(true)

		// comment actions
		case constants.EventComment:
			comment.SetCreated(true)
			comment.SetEdited(true)
		case constants.EventComment + ":" + constants.ActionCreated:
			comment.SetCreated(true)
		case constants.EventComment + ":" + constants.ActionEdited:
			comment.SetEdited(true)

		// schedule actions
		case constants.EventSchedule, constants.EventSchedule + ":run":
			schedule.SetRun(true)
		}
	}

	result.SetPush(push)
	result.SetPullRequest(pull)
	result.SetDeployment(deploy)
	result.SetComment(comment)
	result.SetSchedule(schedule)

	return result
}
