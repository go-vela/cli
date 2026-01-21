// SPDX-License-Identifier: Apache-2.0

package hook

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of build hooks based on the provided configuration.
func (c *Config) Get(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing get for hook configuration")

	// set the pagination options for list of hooks
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	logrus.Tracef("capturing hooks for repo %s/%s", c.Org, c.Repo)

	// send API call to capture a list of hooks
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#HookService.GetAll
	hooks, _, err := client.Hook.GetAll(ctx, c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the hooks in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(hooks)
	case output.DriverJSON:
		// output the hooks in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(hooks, c.Color)
	case output.DriverSpew:
		// output the hooks in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(hooks)
	case "wide":
		// output the hooks in wide table format
		return wideTable(hooks)
	case output.DriverYAML:
		// output the hooks in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(hooks, c.Color)
	default:
		// output the hooks in table format
		return table(hooks)
	}
}
