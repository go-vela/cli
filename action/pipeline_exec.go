// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/compiler/compiler/native"
	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// PipelineExec defines the command for executing a pipeline.
var PipelineExec = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to execute a pipeline.",
	Usage:       "Execute the provided pipeline",
	Action:      pipelineExec,
	Flags: []cli.Flag{

		// Build Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_BRANCH", "PIPELINE_BRANCH"},
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "provide the build branch for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_COMMENT", "PIPELINE_COMMENT"},
			Name:    "comment",
			Aliases: []string{"c"},
			Usage:   "provide the build comment for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_EVENT", "PIPELINE_EVENT"},
			Name:    "event",
			Aliases: []string{"e"},
			Usage:   "provide the build event for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TAG", "PIPELINE_TAG"},
			Name:    "tag",
			Aliases: []string{"t"},
			Usage:   "provide the build tag for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TARGET", "PIPELINE_TARGET"},
			Name:    "target",
			Usage:   "provide the build target for the pipeline",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "PIPELINE_OUTPUT"},
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
		&cli.BoolFlag{
			EnvVars: []string{"VELA_LOCAL", "PIPELINE_LOCAL"},
			Name:    "local",
			Aliases: []string{"l"},
			Usage:   "enables mounting local directory to pipeline",
			Value:   true,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_PATH", "PIPELINE_PATH"},
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "provide the path to the file for the pipeline",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_VOLUMES", "PIPELINE_VOLUMES"},
			Name:    "volume",
			Aliases: []string{"v"},
			Usage:   "provide list of local volumes to mount",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "PIPELINE_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "PIPELINE_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_PIPELINE_TYPE", "PIPELINE_TYPE"},
			Name:    "pipeline-type",
			Aliases: []string{"pt"},
			Usage:   "type of pipeline for the compiler to render",
			Value:   constants.PipelineTypeYAML,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Execute a local Vela pipeline.
    $ {{.HelpName}}
  2. Execute a local Vela pipeline in a nested directory.
    $ {{.HelpName}} --path nested/path/to/dir --file .vela.local.yml
  3. Execute a local Vela pipeline in a specific directory.
    $ {{.HelpName}} --path /absolute/full/path/to/dir --file .vela.local.yml
  4. Execute a local Vela pipeline with ruleset information.
    $ {{.HelpName}} --branch master --event push
  5. Execute a local Vela pipeline with a read-only local volume.
    $ {{.HelpName}} --volume /tmp/foo.txt:/tmp/foo.txt:ro
  6. Execute a local Vela pipeline with a writeable local volume.
    $ {{.HelpName}} --volume /tmp/bar.txt:/tmp/bar.txt:rw
  7. Execute a local Vela pipeline with type of go
    $ {{.HelpName}}  --pipeline-type go

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/exec/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// execute a pipeline.
func pipelineExec(c *cli.Context) error {
	// load variables from the config file
	err := load(c)
	if err != nil {
		return err
	}

	// create the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config
	p := &pipeline.Config{
		Action:       execAction,
		Branch:       c.String("branch"),
		Comment:      c.String("comment"),
		Event:        c.String("event"),
		Tag:          c.String("tag"),
		Target:       c.String("target"),
		Org:          c.String(internal.FlagOrg),
		Repo:         c.String(internal.FlagRepo),
		File:         c.String("file"),
		Local:        c.Bool("local"),
		Path:         c.String("path"),
		Volumes:      c.StringSlice("volume"),
		PipelineType: c.String("pipeline-type"),
	}

	// validate pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Validate
	err = p.Validate()
	if err != nil {
		return err
	}

	// create a compiler client
	//
	// https://godoc.org/github.com/go-vela/compiler/compiler/native#New
	client, err := native.New(c)
	if err != nil {
		return err
	}

	// execute the exec call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Exec
	return p.Exec(client)
}
