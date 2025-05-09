// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating dashboard configuration")

	if c.Action == "add" {
		// check if dashboard name is set
		if len(c.Name) == 0 {
			return fmt.Errorf("no dashboard name provided")
		}
	}

	// check if dashboard action is update or view
	if c.Action == "update" || c.Action == "view" {
		// check if dashboard ID is set
		if len(c.ID) == 0 {
			return fmt.Errorf("no dashboard ID provided")
		}
	}

	// if the dashboard action is update with target repos
	if c.Action == "update" && len(c.TargetRepos) > 0 {
		if len(c.Events) == 0 && len(c.Branches) == 0 {
			return fmt.Errorf("no events or branches updates provided for target repos")
		}
	}

	return nil
}
