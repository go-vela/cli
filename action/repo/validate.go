// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/server/constants"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating repo configuration")

	// check if repository action is not get
	if c.Action != "get" {
		// check if repository org is set
		if len(c.Org) == 0 {
			return fmt.Errorf("no repo org provided")
		}

		// check if repository action is not syncAll
		if c.Action != "syncAll" {
			// check if repository name is set
			if len(c.Name) == 0 {
				return fmt.Errorf("no repo name provided")
			}
		}
	}

	// check if approve build setting is valid if supplied
	if c.Action == "add" || c.Action == "update" {
		if len(c.ApproveBuild) > 0 &&
			c.ApproveBuild != constants.ApproveForkAlways &&
			c.ApproveBuild != constants.ApproveForkNoWrite &&
			c.ApproveBuild != constants.ApproveOnce &&
			c.ApproveBuild != constants.ApproveNever {
			return fmt.Errorf(
				"invalid input for approve-build: must be `%s`, `%s`, `%s`, or `%s`",
				constants.ApproveForkAlways,
				constants.ApproveForkNoWrite,
				constants.ApproveOnce,
				constants.ApproveNever,
			)
		}
	}

	return nil
}
