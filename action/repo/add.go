// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code among actions
package repo

import (
	"fmt"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/library/actions"

	"github.com/sirupsen/logrus"
)

// Add creates a repository based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for repo configuration")

	// create the repository object
	//
	// https://pkg.go.dev/github.com/go-vela/types/library?tab=doc#Repo
	r := &library.Repo{
		Org:          vela.String(c.Org),
		Name:         vela.String(c.Name),
		FullName:     vela.String(fmt.Sprintf("%s/%s", c.Org, c.Name)),
		Link:         vela.String(c.Link),
		Clone:        vela.String(c.Clone),
		Branch:       vela.String(c.Branch),
		BuildLimit:   vela.Int64(c.BuildLimit),
		Timeout:      vela.Int64(c.Timeout),
		Counter:      vela.Int(c.Counter),
		Visibility:   vela.String(c.Visibility),
		Private:      vela.Bool(c.Private),
		Trusted:      vela.Bool(c.Trusted),
		Active:       vela.Bool(c.Active),
		PipelineType: vela.String(c.PipelineType),
		ApproveBuild: vela.String(c.ApproveBuild),
	}

	logrus.Tracef("adding repo %s/%s", c.Org, c.Name)

	if len(c.Events) > 0 {
		populateEvents(r, c.Events)
	}

	// send API call to add a repository
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#RepoService.Add
	repo, _, err := client.Repo.Add(r)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the repository in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(repo)
	case output.DriverJSON:
		// output the repository in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(repo)
	case output.DriverSpew:
		// output the repository in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(repo)
	case output.DriverYAML:
		// output the repository in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(repo)
	default:
		// output the repository in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(repo)
	}
}

// populateEvents is a helper function designed to fill both the legacy `Allow<_>` fields
// as well as the `AllowEvents` struct with the correct values based on a slice input.
func populateEvents(r *library.Repo, events []string) {
	result := new(library.Events)
	push := new(actions.Push)
	pull := new(actions.Pull)
	comment := new(actions.Comment)
	deploy := new(actions.Deploy)

	// -- legacy allow events init --
	r.SetAllowPush(false)
	r.SetAllowPull(false)
	r.SetAllowTag(false)
	r.SetAllowDeploy(false)
	r.SetAllowComment(false)

	// iterate through all events provided
	for _, event := range events {
		switch event {
		case constants.EventPush, constants.EventPush + ":branch":
			r.SetAllowPush(true)
			push.SetBranch(true)
		case constants.EventTag, constants.EventPush + ":" + constants.EventTag:
			r.SetAllowTag(true)
			push.SetTag(true)
		case constants.EventPull, AlternatePull:
			r.SetAllowPull(true)
			pull.SetOpened(true)
			pull.SetReopened(true)
			pull.SetSynchronize(true)
		case constants.EventDeploy, AlternateDeploy, constants.EventDeploy + ":" + constants.ActionCreated:
			r.SetAllowDeploy(true)
			deploy.SetCreated(true)
		case constants.EventComment:
			r.SetAllowComment(true)
			comment.SetCreated(true)
			comment.SetEdited(true)
		case constants.EventDelete:
			push.SetDeleteBranch(true)
			push.SetDeleteTag(true)
		case constants.EventPull + ":" + constants.ActionOpened:
			pull.SetOpened(true)
		case constants.EventPull + ":" + constants.ActionEdited:
			pull.SetEdited(true)
		case constants.EventPull + ":" + constants.ActionSynchronize:
			pull.SetSynchronize(true)
		case constants.EventPull + ":" + constants.ActionReopened:
			pull.SetReopened(true)
		case constants.EventPull + ":" + constants.ActionLabeled:
			pull.SetLabeled(true)
		case constants.EventPull + ":" + constants.ActionUnlabeled:
			pull.SetUnlabeled(true)
		case constants.EventComment + ":" + constants.ActionCreated:
			comment.SetCreated(true)
		case constants.EventComment + ":" + constants.ActionEdited:
			comment.SetEdited(true)
		case constants.EventDelete + ":" + constants.ActionBranch:
			push.SetDeleteBranch(true)
		case constants.EventDelete + ":" + constants.ActionTag:
			push.SetDeleteTag(true)
		}
	}

	result.SetPush(push)
	result.SetPullRequest(pull)
	result.SetDeployment(deploy)
	result.SetComment(comment)

	r.SetAllowEvents(result)
}
