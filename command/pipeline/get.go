// SPDX-License-Identifier: Apache-2.0

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

// CommandGet defines the command for capturing a list of pipelines.
var CommandGet = &cli.Command{
	Name:        "pipeline",
	Aliases:     []string{"pipelines"},
	Description: "Use this command to get a list of pipelines.",
	Usage:       "Display a list of pipelines",
	Action:      get,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "PIPELINE_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the pipeline",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "PIPELINE_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the pipeline",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "PIPELINE_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PAGE", "PIPELINE_PAGE"),
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of pipelines",
			Value:   1,
		},
		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PER_PAGE", "PIPELINE_PER_PAGE"),
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of pipelines to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get pipelines for a repository.
    $ {{.FullName}} --org MyOrg --repo MyRepo
  2. Get pipelines for a repository with wide view output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --output wide
  3. Get pipelines for a repository with yaml output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --output yaml
  4. Get pipelines for a repository with json output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --output json
  5. Get pipelines for a repository when config or environment variables are set.
    $ {{.FullName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of pipelines.
func get(_ context.Context, c *cli.Command) error {
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
		Action:  internal.ActionGet,
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
		Color:   output.ColorOptionsFromCLIContext(c),
	}

	// validate pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Validate
	err = p.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Get
	return p.Get(client)
}
