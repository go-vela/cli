// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// View inspects a dashboard based off the provided configuration.
func (c *Config) View(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing view for dashboard configuration")

	// send API call to capture a dashboard
	dashCard, _, err := client.Dashboard.Get(ctx, c.ID)
	if err != nil {
		return err
	}

	if c.Full {
		err = outputDashboard(dashCard, c)
	} else {
		err = outputDashboard(dashCard.Dashboard, c)
	}

	return err
}

func outputDashboard(dashboard interface{}, c *Config) error {
	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the dashboard in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(dashboard)
	case output.DriverJSON:
		// output the dashboard in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(dashboard, c.Color)
	case output.DriverSpew:
		// output the dashboard in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(dashboard)
	case output.DriverYAML:
		// output the dashboard in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(dashboard, c.Color)
	default:
		// output the dashboard in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(dashboard)
	}
}
