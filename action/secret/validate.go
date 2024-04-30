// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating secret configuration")

	// check if secret file is set
	if len(c.File) > 0 {
		// skip checking all other configuration
		return nil
	}

	// check if secret engine is set
	if len(c.Engine) == 0 {
		return fmt.Errorf("no secret engine provided")
	}

	// check if secret type is set
	if len(c.Type) == 0 {
		return fmt.Errorf("no secret type provided")
	}

	// check if secret org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no secret org provided")
	}

	// check if the secret type provided is valid
	switch c.Type {
	case constants.SecretRepo:
		fallthrough
	case constants.SecretOrg:
		fallthrough
	case constants.SecretShared:
		break
	default:
		return fmt.Errorf("invalid secret type provided: %s", c.Type)
	}

	// check if secret type is repo
	if strings.EqualFold(c.Type, constants.SecretRepo) {
		// check if secret repo is set
		if len(c.Repo) == 0 {
			return fmt.Errorf("no secret repo provided")
		}
	}

	// check if secret type is shared
	if strings.EqualFold(c.Type, constants.SecretShared) {
		// check if secret team is set
		if len(c.Team) == 0 {
			return fmt.Errorf("no secret team provided")
		}
	}

	// check if secret action is remove, update or view
	if c.Action == "remove" || c.Action == "update" || c.Action == "view" {
		// check if secret name is set
		if len(c.Name) == 0 {
			return fmt.Errorf("no secret name provided")
		}
	}

	// check if secret action is add
	if c.Action == "add" {
		// check if secret value is set
		if len(c.Value) == 0 {
			return fmt.Errorf("no secret value provided")
		}
	}

	// check if secret action is add or update
	if c.Action == "add" || c.Action == "update" {
		_, err := library.NewEventsFromSlice(c.AllowEvents)
		if err != nil {
			return err
		}
	}

	return nil
}

// returns a useable list of valid events using a combination of hardcoded shorthand names and AllowEvents.List().
func validEvents() []string {
	unlistedEvents := []string{
		"pull_request",
		"push:branch",
		"push:tag",
		"deployment:created",
		"schedule:run",
	}

	t := true

	evs := library.Events{
		Push: &actions.Push{
			Branch:       &t,
			Tag:          &t,
			DeleteBranch: &t,
			DeleteTag:    &t,
		},
		PullRequest: &actions.Pull{
			Opened:      &t,
			Edited:      &t,
			Synchronize: &t,
			Reopened:    &t,
		},
		Deployment: &actions.Deploy{
			Created: &t,
		},
		Comment: &actions.Comment{
			Created: &t,
			Edited:  &t,
		},
		Schedule: &actions.Schedule{
			Run: &t,
		},
	}

	return append(evs.List(), unlistedEvents...)
}
