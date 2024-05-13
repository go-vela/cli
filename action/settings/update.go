// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/go-vela/cli/internal"
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
	_s, _, err := client.Admin.Settings.Update(s)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(_s)
	case output.DriverJSON:
		// output in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(_s)
	case output.DriverSpew:
		// output in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(_s)
	case output.DriverYAML:
		// output in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(_s)
	default:
		// output in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(_s)
	}
}

// UpdateFromFile updates from a file based on the provided configuration.
func (c *Config) UpdateFromFile(client *vela.Client) error {
	logrus.Debug("executing update from file for platform settings configuration")

	// capture absolute path to file
	path, err := filepath.Abs(c.File)
	if err != nil {
		return err
	}

	logrus.Tracef("reading platform settings contents from %s", path)

	// read contents of file
	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	input := yaml.NewDecoder(bytes.NewReader(contents))

	f := new(ConfigFile)

	for input.Decode(f) == nil {
		if f.Platform == nil {
			return errors.New("invalid input, expected key 'platform'")
		}

		s := &Config{
			Action:            internal.ActionUpdate,
			Output:            c.Output,
			Compiler:          Compiler{},
			Queue:             Queue{},
			RepoAllowlist:     f.Platform.RepoAllowlist,
			ScheduleAllowlist: f.Platform.ScheduleAllowlist,
		}

		// update values if set

		if f.Compiler != nil {
			if f.Compiler.CloneImage != nil {
				s.Compiler.CloneImage = vela.String(f.Compiler.GetCloneImage())
			}

			if f.Compiler.TemplateDepth != nil {
				s.Compiler.TemplateDepth = vela.Int(f.Compiler.GetTemplateDepth())
			}

			if f.Compiler.StarlarkExecLimit != nil {
				s.Compiler.StarlarkExecLimit = vela.UInt64(f.Compiler.GetStarlarkExecLimit())
			}
		}

		if f.Queue != nil {
			if f.Queue.Routes != nil {
				s.Queue.Routes = f.Queue.Routes
			}
		}

		err = s.Validate()
		if err != nil {
			return err
		}

		err = s.Update(client)
		if err != nil {
			return err
		}
	}

	return nil
}
