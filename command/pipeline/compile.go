// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code with expand and view
package pipeline

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandCompile defines the command for compiling a pipeline.
var CommandCompile = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to compile a pipeline.",
	Usage:       "Compile the provided pipeline",
	Action:      compile,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "REPO_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the pipeline",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "REPO_NAME"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the pipeline",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "REPO_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},

		// Pipeline Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REF", "PIPELINE_REF"),
			Name:    "ref",
			Usage:   "provide the repository reference for the pipeline",
			Value:   "main",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Compile a pipeline for a repository.
    $ {{.FullName}} --org MyOrg --repo MyRepo
  2. Compile a pipeline for a repository with json output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --output json
  3. Compile a pipeline for a repository when config or environment variables are set.
    $ {{.FullName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/compile/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to compile a pipeline.
func compile(_ context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
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
		Action: internal.ActionCompile,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
		Ref:    c.String("ref"),
	}

	// validate pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Validate
	err = p.Validate()
	if err != nil {
		return err
	}

	// execute the compile call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Compile
	return p.Compile(client)
}
