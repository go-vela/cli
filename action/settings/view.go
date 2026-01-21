// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// View inspects settings based off the provided configuration.
func (c *Config) View(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing view for settings configuration")

	logrus.Trace("inspecting settings")

	response, _, err := client.Admin.Settings.Get(ctx)
	if err != nil {
		return fmt.Errorf("unable to retrieve settings: %w", err)
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(response)
	case output.DriverJSON:
		// output in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(response, c.Color)
	case output.DriverSpew:
		// output in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(response)
	case output.DriverYAML:
		// output in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(response, c.Color)
	default:
		// output in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(response)
	}
}
