// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/pipeline"

	"github.com/go-vela/compiler/compiler/native"

	"github.com/urfave/cli/v2"
)

// PipelineValidate defines the command for verifying a pipeline.
var PipelineValidate = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to validate a pipeline.",
	Usage:       "Validate a Vela pipeline",
	Action:      pipelineValidate,
	Flags: []cli.Flag{

		// Pipeline Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_FILE", "PIPELINE_FILE"},
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "provide the file name for the pipeline",
			Value:   ".vela.yml",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_PATH", "PIPELINE_PATH"},
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "provide the path to the file for the pipeline",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Validate a Vela pipeline.
    $ {{.HelpName}}
  2. Validate a Vela pipeline in a nested directory.
    $ {{.HelpName}} --path nested/path/to/dir
  3. Validate a Vela pipeline in a specific directory.
    $ {{.HelpName}} --path /absolute/full/path/to/dir

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/pipeline/validate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// verify a pipeline.
func pipelineValidate(c *cli.Context) error {
	// create a compiler client
	//
	// https://godoc.org/github.com/go-vela/compiler/compiler/native#New
	client, err := native.New(c)
	if err != nil {
		return err
	}

	// create the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config
	p := &pipeline.Config{
		Action: validateAction,
		File:   c.String("file"),
		Path:   c.String("path"),
	}

	// validate pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Validate
	err = p.Validate()
	if err != nil {
		return err
	}

	// execute the validate file call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.ValidateFile
	return p.ValidateFile(client)
}
