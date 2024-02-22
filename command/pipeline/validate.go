// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/types/constants"
	"github.com/sirupsen/logrus"

	"github.com/go-vela/server/compiler/native"
	"github.com/go-vela/server/util"

	"github.com/urfave/cli/v2"
)

// CommandValidate defines the command for verifying a pipeline.
var CommandValidate = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to validate a pipeline.",
	Usage:       "Validate a Vela pipeline",
	Action:      validate,
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
			Value:   "main",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_TEMPLATE", "PIPELINE_TEMPLATE"},
			Name:    "template",
			Usage:   "DEPRECATED (Vela CLI will attempt to fetch templates if they exist)",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_TEMPLATE_FILE", "PIPELINE_TEMPLATE_FILE"},
			Name:    "template-file",
			Usage:   "enables using a local template file for expansion",
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_MAX_TEMPLATE_DEPTH", "MAX_TEMPLATE_DEPTH"},
			Name:    "max-template-depth",
			Usage:   "set the maximum depth for nested templates",
			Value:   3,
		},
		&cli.Uint64Flag{
			EnvVars: []string{"VELA_COMPILER_STARLARK_EXEC_LIMIT", "COMPILER_STARLARK_EXEC_LIMIT"},
			Name:    "compiler-starlark-exec-limit",
			Aliases: []string{"starlark-exec-limit", "sel"},
			Usage:   "set the starlark execution step limit for compiling starlark pipelines",
			Value:   7500,
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
  6. Validate a template pipeline with expanding steps (when templates are sourced from private Github instance)
    $ {{.HelpName}} --compiler.github.token <token> --compiler.github.url <url>
  7. Validate a local template pipeline with expanding steps
    $ {{.HelpName}} --template-file name:/path/to/file
  8. Validate a local, nested template pipeline with custom template depth.
    $ {{.HelpName}} --template-file name:/path/to/file name:/path/to/file --max-template-depth 2
DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/validate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to verify a pipeline.
func validate(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	if c.Bool("template") {
		logrus.Warnf("`template` flag is deprecated and will be removed in a later release")
	}

	// create the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config
	p := &pipeline.Config{
		Action:        internal.ActionValidate,
		Org:           c.String(internal.FlagOrg),
		Repo:          c.String(internal.FlagRepo),
		File:          c.String("file"),
		Path:          c.String("path"),
		Ref:           c.String("ref"),
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
	// https://godoc.org/github.com/go-vela/server/compiler/native#New
	client, err := native.New(c)
	if err != nil {
		return err
	}

	// set starlark exec limit
	client.StarlarkExecLimit = c.Uint64("compiler-starlark-exec-limit")

	// set when user is sourcing templates from local machine
	if len(p.TemplateFiles) != 0 {
		client.WithLocalTemplates(p.TemplateFiles)
		client.TemplateDepth = c.Int("max-template-depth")
	} else {
		// set max template depth to minimum of 5 and provided value if local templates are not provided.
		// This prevents users from spamming SCM
		client.TemplateDepth = util.MinInt(c.Int("max-template-depth"), 5)
		logrus.Debugf("no local template files provided, setting max template depth to %d", client.TemplateDepth)
	}

	// execute the validate local call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.ValidateLocal
	return p.ValidateLocal(client.WithLocal(true).WithPrivateGitHub(c.String(internal.FlagCompilerGitHubURL), c.String(internal.FlagCompilerGitHubToken)))
}
