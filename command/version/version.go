// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package version

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/cli/version"

	"github.com/urfave/cli/v2"
)

// CommandVersion defines the command for returning version information on the CLI.
var CommandVersion = &cli.Command{
	Name:        "version",
	Description: "Use this command to output version information.",
	Usage:       "Output version information",
	Action:      runVersion,
	Flags: []cli.Flag{

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "STEP_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Output Vela version information
    $ {{.HelpName}} --api.addr https://vela.example.com
  2. Output Vela version information with JSON output
    $ {{.HelpName}} --output json

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/version/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to output version
// information.
func runVersion(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.String(internal.FlagOutput) {
	case output.DriverDump:
		// output the version in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(version.New())
	case output.DriverJSON:
		// output the version in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(version.New())
	case output.DriverSpew:
		// output the version in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(version.New())
	case output.DriverYAML:
		// output the version in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(version.New())
	default:
		// output the version in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(version.New())
	}
}
