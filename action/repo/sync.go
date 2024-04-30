// SPDX-License-Identifier: Apache-2.0
package repo

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Sync synchronizes a single repository in the Vela Database with the SCM.
func (c *Config) Sync(client *vela.Client) error {
	logrus.Debug("executing SCM sync for repo")

	logrus.Tracef("syncing repo %s/%s", c.Org, c.Name)

	// send API call to sync repository
	msg, _, err := client.SCM.Sync(c.Org, c.Name)
	if err != nil {
		return err
	}

	return output.Stdout(*msg)
}

// SyncAll synchronizes all org repositories in the Vela Database with the SCM.
func (c *Config) SyncAll(client *vela.Client) error {
	logrus.Debug("executing SCM sync for org repos")

	logrus.Tracef("syncing repos for org: %s...", c.Org)

	// send API call to sync org repos
	msg, _, err := client.SCM.SyncAll(c.Org)
	if err != nil {
		return err
	}

	return output.Stdout(*msg)
}
