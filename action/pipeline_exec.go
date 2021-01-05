// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/internal"

	"github.com/urfave/cli/v2"
)

// PipelineExec defines the command for executing a pipeline.
var PipelineExec = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to execute a pipeline.",
	Usage:       "Execute the provided pipeline",
	Action:      pipelineExec,
	Flags: []cli.Flag{

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "REPO_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},

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
  1. Execute a local Vela pipeline.
    $ {{.HelpName}}
  2. Execute a local Vela pipeline in a nested directory.
    $ {{.HelpName}} --path nested/path/to/dir
  3. Execute a local Vela pipeline in a specific directory.
    $ {{.HelpName}} --path /absolute/full/path/to/dir

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/pipeline/exec/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// execute a pipeline.
func pipelineExec(c *cli.Context) error {
	// TODO: implement in a future PR
	//
	// currently does nothing to keep PR sizes smaller
	return nil
}
