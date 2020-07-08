// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"fmt"
	"strings"

	"github.com/go-vela/types/constants"

	"github.com/sirupsen/logrus"
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

	// check if secret action is add or update
	if c.Action == "add" || c.Action == "update" {
		// check if secret value is set
		if len(c.Value) == 0 {
			return fmt.Errorf("no secret value provided")
		}
	}

	return nil
}
