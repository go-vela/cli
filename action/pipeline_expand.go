// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// PipelineExpand defines the command for expanding a pipeline.
var PipelineExpand = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to expand a pipeline.",
	Usage:       "Expand the provided pipeline",
	Action:      pipelineExpand,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "REPO_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "REPO_NAME"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the pipeline",
		},

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
			EnvVars: []string{"VELA_REF", "PIPELINE_REF"},
			Name:    "ref",
			Usage:   "provide the repository reference for the pipeline",
			Value:   "master",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Expand a pipeline for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo
  2. Expand a pipeline for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output json
  3. Expand a pipeline for a repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/expand/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// expand a pipeline.
func pipelineExpand(c *cli.Context) error {
	// load variables from the config file
	err := load(c)
	if err != nil {
		return err
	}

	// parse the Vela client from the context
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/client?tab=doc#Parse
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// create the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config
	p := &pipeline.Config{
		Action: expandAction,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Output: c.String(internal.FlagOutput),
		Ref:    c.String("ref"),
	}

	// validate pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Validate
	err = p.Validate()
	if err != nil {
		return err
	}

	// execute the expand call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Expand
	return p.Expand(client)
}
