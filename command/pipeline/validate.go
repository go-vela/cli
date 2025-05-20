// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/server/compiler/native"
	"github.com/go-vela/server/constants"
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
			Sources:  cli.EnvVars("VELA_ORG", "REPO_ORG"),
			Name:     internal.FlagOrg,
			Aliases:  []string{"o"},
			Usage:    "provide the organization for the pipeline",
			Category: "1. Repo:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_REPO", "REPO_NAME"),
			Name:     internal.FlagRepo,
			Aliases:  []string{"r"},
			Usage:    "provide the repository for the pipeline",
			Category: "1. Repo:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_PIPELINE_TYPE", "PIPELINE_TYPE"),
			Name:     "pipeline-type",
			Aliases:  []string{"pt"},
			Usage:    "type of pipeline for the compiler to render",
			Value:    constants.PipelineTypeYAML,
			Category: "1. Repo:",
		},

		// Pipeline Flags

		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_FILE", "PIPELINE_FILE"),
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "provide the file name for the pipeline",
			Value:    ".vela.yml",
			Category: "2. Pipeline:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_PATH", "PIPELINE_PATH"),
			Name:     "path",
			Aliases:  []string{"p"},
			Usage:    "provide the path to the file for the pipeline",
			Category: "2. Pipeline:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_REF", "PIPELINE_REF"),
			Name:     "ref",
			Usage:    "provide the repository reference for the pipeline",
			Value:    "main",
			Category: "2. Pipeline:",
		},
		&cli.BoolFlag{
			Sources:  cli.EnvVars("VELA_TEMPLATE", "PIPELINE_TEMPLATE"),
			Name:     "template",
			Usage:    "DEPRECATED (Vela CLI will attempt to fetch templates if they exist)",
			Category: "2. Pipeline:",
		},
		&cli.StringSliceFlag{
			Sources:  cli.EnvVars("VELA_TEMPLATE_FILE", "PIPELINE_TEMPLATE_FILE"),
			Name:     "template-file",
			Usage:    "enables using a local template file for expansion",
			Category: "2. Pipeline:",
		},
		&cli.IntFlag{
			Sources:  cli.EnvVars("VELA_MAX_TEMPLATE_DEPTH", "MAX_TEMPLATE_DEPTH"),
			Name:     "max-template-depth",
			Usage:    "set the maximum depth for nested templates",
			Value:    3,
			Category: "2. Pipeline:",
		},
		&cli.Int64Flag{
			Sources:  cli.EnvVars("VELA_COMPILER_STARLARK_EXEC_LIMIT", "COMPILER_STARLARK_EXEC_LIMIT"),
			Name:     "compiler-starlark-exec-limit",
			Aliases:  []string{"starlark-exec-limit", "sel"},
			Usage:    "set the starlark execution step limit for compiling starlark pipelines",
			Value:    7500,
			Category: "2. Pipeline:",
		},
		&cli.BoolFlag{
			Sources:  cli.EnvVars("VELA_REMOTE", "PIPELINE_REMOTE"),
			Name:     "remote",
			Usage:    "enables validating a pipeline on a remote server",
			Value:    false,
			Category: "2. Pipeline:",
		},

		// RuleData Flags
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_BRANCH", "PIPELINE_BRANCH", "VELA_BUILD_BRANCH"),
			Name:     "branch",
			Aliases:  []string{"b"},
			Usage:    "provide the build branch for the pipeline",
			Category: "3. Ruleset:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_COMMENT", "PIPELINE_COMMENT", "VELA_BUILD_COMMENT"),
			Name:     "comment",
			Aliases:  []string{"c"},
			Usage:    "provide the build comment for the pipeline",
			Category: "3. Ruleset:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_EVENT", "PIPELINE_EVENT", "VELA_BUILD_EVENT"),
			Name:     "event",
			Aliases:  []string{"e"},
			Usage:    "provide the build event for the pipeline",
			Category: "3. Ruleset:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_STATUS", "PIPELINE_STATUS", "VELA_BUILD_STATUS"),
			Name:     "status",
			Usage:    "provide the expected build status for the local validation",
			Value:    "success",
			Category: "3. Ruleset:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_TAG", "PIPELINE_TAG", "VELA_BUILD_TAG"),
			Name:     "tag",
			Usage:    "provide the build tag for the pipeline",
			Category: "3. Ruleset:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_TARGET", "PIPELINE_TARGET", "VELA_BUILD_TARGET"),
			Name:     "target",
			Usage:    "provide the build target for the pipeline",
			Category: "3. Ruleset:",
		},
		&cli.StringSliceFlag{
			Sources:  cli.EnvVars("VELA_FILE_CHANGESET", "FILE_CHANGESET"),
			Name:     "file-changeset",
			Aliases:  []string{"fcs"},
			Usage:    "provide a list of files changed for ruleset matching",
			Category: "3. Ruleset:",
		},

		// Compiler Flags

		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_COMPILER_GITHUB_TOKEN", "COMPILER_GITHUB_TOKEN"),
			Name:     internal.FlagCompilerGitHubToken,
			Aliases:  []string{"ct"},
			Usage:    "github compiler token",
			Category: "4. Compiler:",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_COMPILER_GITHUB_URL", "COMPILER_GITHUB_URL"),
			Name:     internal.FlagCompilerGitHubURL,
			Aliases:  []string{"cgu"},
			Usage:    "github url, used by compiler, for pulling registry templates",
			Category: "4. Compiler:",
			Value:    "https://github.com",
		},
		&cli.StringFlag{
			Sources:  cli.EnvVars("VELA_CLONE_IMAGE", "COMPILER_CLONE_IMAGE"),
			Name:     "clone-image",
			Usage:    "the clone image to use for the injected clone step",
			Value:    "docker.io/target/vela-git-slim:v0.12.0@sha256:533786ab3ef17c900b0105fdffbd7143d2601803f28b39e156132ad25819af0f",
			Category: "4. Compiler:",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Validate a local Vela pipeline.
    $ {{.FullName}}
  2. Validate a local Vela pipeline in a nested directory.
    $ {{.FullName}} --path nested/path/to/dir
  3. Validate a local Vela pipeline in a specific directory.
    $ {{.FullName}} --path /absolute/full/path/to/dir
  4. Validate a remote pipeline for a repository.
    $ {{.FullName}} --remote --org MyOrg --repo MyRepo
  5. Validate a remote pipeline for a repository with json output.
    $ {{.FullName}} --remote --org MyOrg --repo MyRepo --output json
  6. Validate a template pipeline with expanding steps (when templates are sourced from private Github instance)
    $ {{.FullName}} --compiler.github.token <token> --compiler.github.url <url>
  7. Validate a pipeline with ruleset data
    $ {{.FullName}} --branch dev --event push
  8. Validate a local template pipeline with expanding steps
    $ {{.FullName}} --template-file name:/path/to/file
  9. Validate a local, nested template pipeline with custom template depth.
    $ {{.FullName}} --template-file name:/path/to/file name:/path/to/file --max-template-depth 2
DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/validate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to verify a pipeline.
func validate(ctx context.Context, c *cli.Command) error {
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
		Branch:        c.String("branch"),
		Comment:       c.String("comment"),
		Event:         c.String("event"),
		FileChangeset: c.StringSlice("file-changeset"),
		Status:        c.String("status"),
		Tag:           c.String("tag"),
		Target:        c.String("target"),
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
	client, err := native.FromCLICommand(ctx, c)
	if err != nil {
		return err
	}

	// set starlark exec limit
	client.SetStarlarkExecLimit(c.Int64("compiler-starlark-exec-limit"))

	// set when user is sourcing templates from local machine
	if len(p.TemplateFiles) != 0 {
		client.WithLocalTemplates(p.TemplateFiles)
		client.SetTemplateDepth(c.Int("max-template-depth"))
	} else {
		// set max template depth to minimum of 5 and provided value if local templates are not provided.
		// This prevents users from spamming SCM
		client.SetTemplateDepth(min(c.Int("max-template-depth"), 5))
		logrus.Debugf("no local template files provided, setting max template depth to %d", client.GetTemplateDepth())
	}

	// execute the validate local call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.ValidateLocal
	//nolint:contextcheck // consider refactor to add context to action
	return p.ValidateLocal(client.WithLocal(true).WithPrivateGitHub(ctx, c.String(internal.FlagCompilerGitHubURL), c.String(internal.FlagCompilerGitHubToken)))
}
