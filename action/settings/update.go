// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/api/types/settings"
)

// Update modifies settings based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
	logrus.Debug("executing update for settings configuration")

	// create the settings object
	s := &settings.Platform{
		Queue: &settings.Queue{
			Routes: c.Queue.Routes,
		},
		Compiler: &settings.Compiler{
			CloneImage:        c.Compiler.CloneImage,
			TemplateDepth:     c.Compiler.TemplateDepth,
			StarlarkExecLimit: c.Compiler.StarlarkExecLimit,
		},
		RepoAllowlist:     c.RepoAllowlist,
		ScheduleAllowlist: c.ScheduleAllowlist,
	}

	logrus.Trace("updating settings")

	// send API call to modify settings
	s_, _, err := client.Admin.Settings.Update(s)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(s_)
	case output.DriverJSON:
		// output in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(s_)
	case output.DriverSpew:
		// output in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(s_)
	case output.DriverYAML:
		// output in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(s_)
	default:
		// output in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(s_)
	}
}
