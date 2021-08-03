// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/types/constants"

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
		&cli.StringFlag{
			EnvVars: []string{"VELA_PIPELINE_TYPE", "PIPELINE_TYPE"},
			Name:    "pipeline-type",
			Aliases: []string{"pt"},
			Usage:   "type of pipeline for the compiler to render",
			Value:   constants.PipelineTypeYAML,
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
		&cli.StringFlag{
			EnvVars: []string{"VELA_REF", "PIPELINE_REF"},
			Name:    "ref",
			Usage:   "provide the repository reference for the pipeline",
			Value:   "master",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_TEMPLATE", "PIPELINE_TEMPLATE"},
			Name:    "template",
			Usage:   "enables validating a pipeline with templates",
			Value:   false,
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_TEMPLATE_FILE", "PIPELINE_TEMPLATE_FILE"},
			Name:    "template-file",
			Usage:   "enables using a local template file for expansion",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_REMOTE", "PIPELINE_REMOTE"},
			Name:    "remote",
			Usage:   "enables validating a pipeline on a remote server",
			Value:   false,
		},

		// Compiler Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_COMPILER_GITHUB_TOKEN", "COMPILER_GITHUB_TOKEN"},
			Name:    internal.FlagCompilerGitHubToken,
			Aliases: []string{"ct"},
			Usage:   "github compiler token",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_COMPILER_GITHUB_URL", "COMPILER_GITHUB_URL"},
			Name:    internal.FlagCompilerGitHubURL,
			Aliases: []string{"cgu"},
			Usage:   "github url, used by compiler, for pulling registry templates",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Validate a local Vela pipeline.
    $ {{.HelpName}}
  2. Validate a local Vela pipeline in a nested directory.
    $ {{.HelpName}} --path nested/path/to/dir
  3. Validate a local Vela pipeline in a specific directory.
    $ {{.HelpName}} --path /absolute/full/path/to/dir
  4. Validate a remote pipeline for a repository.
    $ {{.HelpName}} --remote --org MyOrg --repo MyRepo
  5. Validate a remote pipeline for a repository with json output.
    $ {{.HelpName}} --remote --org MyOrg --repo MyRepo --output json
  5. Validate a template pipeline with expanding steps
    $ {{.HelpName}} --template
  6. Validate a local template pipeline with expanding steps
    $ {{.HelpName}} --template --template-file name:/path/to/file
DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/validate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// verify a pipeline.
func pipelineValidate(c *cli.Context) error {
	// load variables from the config file
	err := load(c)
	if err != nil {
		return err
	}

	// create the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config
	p := &pipeline.Config{
		Action:        validateAction,
		Org:           c.String(internal.FlagOrg),
		Repo:          c.String(internal.FlagRepo),
		File:          c.String("file"),
		Path:          c.String("path"),
		Ref:           c.String("ref"),
		Template:      c.Bool("template"),
		TemplateFiles: c.StringSlice("template-file"),
		Remote:        c.Bool("remote"),
		PipelineType:  c.String("pipeline-type"),
	}

	// validate pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Validate
	err = p.Validate()
	if err != nil {
		return err
	}

	// check if pipeline org is provided
	if len(p.Org) > 0 && len(p.Repo) > 0 && p.Remote {
		// parse the Vela client from the context
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/client?tab=doc#Parse
		client, err := client.Parse(c)
		if err != nil {
			return err
		}

		// execute the validate remote call for the pipeline configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.ValidateRemote
		return p.ValidateRemote(client)
	}

	// create a compiler client
	//
	// https://godoc.org/github.com/go-vela/compiler/compiler/native#New
	client, err := native.New(c)
	if err != nil {
		return err
	}

	// set when user is sourcing templates from local machine
	if len(p.TemplateFiles) != 0 {
		client.WithLocal(true)
	}

	// execute the validate local call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.ValidateLocal
	return p.ValidateLocal(client.WithPrivateGitHub(c.String(internal.FlagCompilerGitHubURL), c.String(internal.FlagCompilerGitHubToken)))
}
