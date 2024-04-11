// SPDX-License-Identifier: Apache-2.0

package build

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Approve approves a build based off the provided configuration.
func (c *Config) Approve(client *vela.Client) error {
	logrus.Debug("executing approve for build configuration")

	logrus.Tracef("approving build %s/%s/%d", c.Org, c.Repo, c.Number)

	// send API call to approve a build
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#BuildService.Approve
	_, err := client.Build.Approve(c.Org, c.Repo, c.Number)
	if err != nil {
		return err
	}

	// output the build in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(fmt.Sprintf("successfully approved build %s/%s/%d", c.Org, c.Repo, c.Number))
}
