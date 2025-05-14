// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
)

// CommandGenerate defines the command for producing a pipeline.
var CommandGenerate = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to generate a pipeline.",
	Usage:       "Generate a valid Vela pipeline",
	Action:      generate,
	Flags: []cli.Flag{

		// Pipeline Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_FILE", "PIPELINE_FILE"),
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "provide the file name for the pipeline",
			Value:   ".vela.yml",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_PATH", "PIPELINE_PATH"),
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "provide the path to the file for the pipeline",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_STAGES", "PIPELINE_STAGES"),
			Name:    "stages",
			Aliases: []string{"s"},
			Usage:   "enable generating the pipeline with stages",
			Value:   "false",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_TYPE", "PIPELINE_TYPE"),
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   "provide the type of pipeline being generated",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Generate a Vela pipeline.
    $ {{.HelpName}}
  2. Generate a Vela pipeline in a nested directory.
    $ {{.HelpName}} --path nested/path/to/dir
  3. Generate a Vela pipeline in a specific directory.
    $ {{.HelpName}} --path /absolute/full/path/to/dir
  4. Generate a Vela pipeline with stages.
    $ {{.HelpName}} --stages true
  5. Generate a go Vela pipeline.
    $ {{.HelpName}} --secret.type go
  6. Generate a java Vela pipeline.
    $ {{.HelpName}} --secret.type java
  7. Generate a node Vela pipeline.
    $ {{.HelpName}} --secret.type node

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/generate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to produce a pipeline.
func generate(_ context.Context, c *cli.Command) error {
	// create the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config
	p := &pipeline.Config{
		Action: internal.ActionGenerate,
		File:   c.String("file"),
		Path:   c.String("path"),
		Stages: c.Bool("stages"),
		Type:   c.String("type"),
	}

	// validate pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Validate
	err := p.Validate()
	if err != nil {
		return err
	}

	// execute the generate call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Generate
	return p.Generate()
}
