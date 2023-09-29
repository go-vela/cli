// SPDX-License-Identifier: Apache-2.0

package hook

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// View inspects a hook based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	logrus.Debug("executing view for hook configuration")

	logrus.Tracef("inspecting hook %s/%s/%d", c.Org, c.Repo, c.Number)

	// send API call to capture a hook
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#HookService.Get
	hook, _, err := client.Hook.Get(c.Org, c.Repo, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the hook in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(hook)
	case output.DriverJSON:
		// output the hook in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(hook)
	case output.DriverSpew:
		// output the hook in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(hook)
	case output.DriverYAML:
		// output the hook in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(hook)
	default:
		// output the hook in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(hook)
	}
}
