// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/pipeline"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/server/compiler/native"
	"github.com/go-vela/server/util"
	"github.com/go-vela/types/constants"
)

// CommandExec defines the command for executing a pipeline.
var CommandExec = &cli.Command{
	Name:        "pipeline",
	Description: "Use this command to execute a pipeline.",
	Usage:       "Execute the provided pipeline",
	Action:      exec,
	Flags: []cli.Flag{

		// Build Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_BRANCH", "PIPELINE_BRANCH", "VELA_BUILD_BRANCH"},
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "provide the build branch for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_COMMENT", "PIPELINE_COMMENT", "VELA_BUILD_COMMENT"},
			Name:    "comment",
			Aliases: []string{"c"},
			Usage:   "provide the build comment for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_EVENT", "PIPELINE_EVENT", "VELA_BUILD_EVENT"},
			Name:    "event",
			Aliases: []string{"e"},
			Usage:   "provide the build event for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TAG", "PIPELINE_TAG", "VELA_BUILD_TAG"},
			Name:    "tag",
			Usage:   "provide the build tag for the pipeline",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TARGET", "PIPELINE_TARGET", "VELA_BUILD_TARGET"},
			Name:    "target",
			Usage:   "provide the build target for the pipeline",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_FILE_CHANGESET", "FILE_CHANGESET"},
			Name:    "file-changeset",
			Aliases: []string{"fcs"},
			Usage:   "provide a list of files changed for ruleset matching",
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

		// Runtime Flags

		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_VOLUMES", "PIPELINE_VOLUMES"},
			Name:    "volume",
			Aliases: []string{"v"},
			Usage:   "provide list of local volumes to mount",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_PRIVILEGED_IMAGES", "PIPELINE_PRIVILEGED_IMAGES"},
			Name:    "privileged-images",
			Aliases: []string{"pi"},
			Usage:   "provide list of pipeline images that will run in privileged mode",
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

		// Compiler Template Flags

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
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_TEMPLATE_FILE", "PIPELINE_TEMPLATE_FILE"},
			Name:    "template-file",
			Aliases: []string{"tf, tfs, template-files"},
			Usage:   "enables using a local template file for expansion in the form <name>:<path>",
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

		// Environment Flags
		&cli.BoolFlag{
			EnvVars: []string{"VELA_ENV_FILE", "ENV_FILE"},
			Name:    "env-file",
			Aliases: []string{"ef"},
			Usage:   "load environment variables from a .env file",
			Value:   false,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ENV_FILE_PATH", "ENV_FILE_PATH"},
			Name:    "env-file-path",
			Aliases: []string{"efp"},
			Usage:   "provide the path to the file for the environment",
		},
		&cli.BoolFlag{
			EnvVars: []string{"ONBOARD_LOCAL_ENV", "LOCAL_ENV"},
			Name:    "local-env",
			Usage:   "load environment variables from local environment",
			Value:   false,
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_ENV_VARS"},
			Name:    "env-vars",
			Aliases: []string{"env"},
			Usage:   "load a set of environment variables in the form of KEY1=VAL1,KEY2=VAL2",
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
    $ {{.HelpName}} --branch main --event push
  5. Execute a local Vela pipeline with a read-only local volume.
    $ {{.HelpName}} --volume /tmp/foo.txt:/tmp/foo.txt:ro
  6. Execute a local Vela pipeline with a writeable local volume.
    $ {{.HelpName}} --volume /tmp/bar.txt:/tmp/bar.txt:rw
  7. Execute a local Vela pipeline with type of go
    $ {{.HelpName}} --pipeline-type go
  8. Execute a local Vela pipeline with local templates
    $ {{.HelpName}} --template-file <template_name>:<path_to_template>
  9. Execute a local Vela pipeline with specific environment variables
    $ {{.HelpName}} --env KEY1=VAL1,KEY2=VAL2
  10. Execute a local Vela pipeline with your existing local environment loaded into pipeline
    $ {{.HelpName}} --local-env
  11. Execute a local Vela pipeline with an environment file loaded in
    $ {{.HelpName}} --env-file (uses .env by default)
      OR
    $ {{.HelpName}} --env-file-path <path_to_file>
  12. Execute a local Vela pipeline using remote templates
    $ {{.HelpName}} --compiler.github.token <GITHUB_PAT> --compiler.github.url <GITHUB_URL>

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/pipeline/exec/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to execute a pipeline.
func exec(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// clear local environment unless told otherwise
	if !c.Bool("local-env") {
		os.Clearenv()
	}

	// iterate through command-based env variables and set them in environment
	for _, envSet := range c.StringSlice("env-vars") {
		parts := strings.SplitN(envSet, "=", 2)

		os.Setenv(parts[0], parts[1])
	}

	// load env file if provided
	if c.Bool("env-file") || len(c.String("env-file-path")) > 0 {
		switch len(c.String("env-file-path")) {
		case 0:
			err := godotenv.Load()
			if err != nil {
				logrus.Fatal("Error loading env file")
			}
		default:
			err := godotenv.Load(c.String("env-file-path"))
			if err != nil {
				logrus.Fatal("Error loading env file")
			}
		}
	}

	// account for users omitting the `refs/tags` prefix of the tag value
	tag := c.String("tag")

	if len(tag) > 0 {
		if !strings.HasPrefix(tag, "refs/tags/") {
			logrus.Debugf("setting tag value to refs/tags/%s", tag)
			tag = "refs/tags/" + tag
		}
	}

	// create the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config
	p := &pipeline.Config{
		Action:           internal.ActionExec,
		Branch:           c.String("branch"),
		Comment:          c.String("comment"),
		Event:            c.String("event"),
		Tag:              tag,
		Target:           c.String("target"),
		Org:              c.String(internal.FlagOrg),
		Repo:             c.String(internal.FlagRepo),
		File:             c.String("file"),
		FileChangeset:    c.StringSlice("file-changeset"),
		TemplateFiles:    c.StringSlice("template-file"),
		Local:            c.Bool("local"),
		Path:             c.String("path"),
		Volumes:          c.StringSlice("volume"),
		PrivilegedImages: c.StringSlice("privileged-images"),
		PipelineType:     c.String("pipeline-type"),
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
	// https://godoc.org/github.com/go-vela/server/compiler/native#New
	client, err := native.New(c)
	if err != nil {
		return err
	}

	// set starlark exec limit
	client.SetStarlarkExecLimit(c.Uint64("compiler-starlark-exec-limit"))

	// set when user is sourcing templates from local machine
	if len(p.TemplateFiles) != 0 {
		client.WithLocalTemplates(p.TemplateFiles)
		client.SetTemplateDepth(util.MinInt(c.Int("max-template-depth"), 10))
	} else {
		// set max template depth to minimum of 5 and provided value if local templates are not provided.
		// This prevents users from spamming SCM
		client.SetTemplateDepth(util.MinInt(c.Int("max-template-depth"), 5))
		logrus.Debugf("no local template files provided, setting max template depth to %d", client.TemplateDepth)
	}

	// execute the exec call for the pipeline configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/pipeline?tab=doc#Config.Exec
	return p.Exec(client.WithPrivateGitHub(c.String(internal.FlagCompilerGitHubURL), c.String(internal.FlagCompilerGitHubToken)))
}
