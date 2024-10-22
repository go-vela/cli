// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/constants"
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
		_, err := api.NewEventsFromSlice(c.AllowEvents)
		if err != nil {
			return err
		}
	}

	return nil
}
