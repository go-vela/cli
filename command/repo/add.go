// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore duplicate of update
package repo

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/server/constants"
)

// CommandAdd defines the command for creating a repository.
var CommandAdd = &cli.Command{
	Name:        "repo",
	Description: "Use this command to add a repository.",
	Usage:       "Add a new repository from the provided configuration",
	Action:      add,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "REPO_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the repository",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "REPO_NAME"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the name for the repository",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_BRANCH", "REPO_BRANCH"),
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "default branch for the repository",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_LINK", "REPO_LINK"),
			Name:    "link",
			Aliases: []string{"l"},
			Usage:   "full URL to repository in source control",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_CLONE", "REPO_CLONE"),
			Name:    "clone",
			Aliases: []string{"c"},
			Usage:   "full clone URL to repository in source control",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_VISIBILITY", "REPO_VISIBILITY"),
			Name:    "visibility",
			Aliases: []string{"v"},
			Usage:   "access level required to view the repository",
			Value:   constants.VisibilityPublic,
		},
		&cli.Int32Flag{
			Sources: cli.EnvVars("VELA_BUILD_LIMIT", "REPO_BUILD_LIMIT"),
			Name:    "build.limit",
			Usage:   "limit of concurrent builds allowed in repository",
			Value:   constants.BuildLimitDefault,
		},
		&cli.Int32Flag{
			Sources: cli.EnvVars("VELA_TIMEOUT", "REPO_TIMEOUT"),
			Name:    "timeout",
			Aliases: []string{"t"},
			Usage:   "max time allowed per build in repository",
			Value:   constants.BuildTimeoutDefault,
		},
		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_COUNTER", "REPO_COUNTER"),
			Name:    "counter",
			Aliases: []string{"ct"},
			Usage:   "set a value for a new build number",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_PRIVATE", "REPO_PRIVATE"),
			Name:    "private",
			Aliases: []string{"p"},
			Usage:   "disable public access to the repository",
			Value:   "false",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_TRUSTED", "REPO_TRUSTED"),
			Name:    "trusted",
			Aliases: []string{"tr"},
			Usage:   "elevated permissions for builds executed for repo",
			Value:   "false",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ACTIVE", "REPO_ACTIVE"),
			Name:    "active",
			Aliases: []string{"a"},
			Usage:   "current status of the repository",
			Value:   "true",
		},
		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_EVENTS", "REPO_EVENTS", "VELA_ADD_EVENTS", "REPO_ADD_EVENTS"),
			Name:    "event",
			Aliases: []string{"events", "e"},
			Usage:   "webhook event(s) repository responds to",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_PIPELINE_TYPE", "PIPELINE_TYPE"),
			Name:    "pipeline-type",
			Aliases: []string{"pt"},
			Usage:   "type of base pipeline for the compiler to render",
			Value:   constants.PipelineTypeYAML,
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_APPROVE_BUILD", "REPO_APPROVE_BUILD"),
			Name:    "approve-build",
			Aliases: []string{"ab", "approve-build-setting"},
			Usage:   "when to require admin approval to run builds from outside contributors (`fork-always`, `fork-no-write`, or `never`)",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "REPO_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Add a repository with push and pull request enabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --event push --event pull_request
  2. Add a repository with all event types enabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --event push --event pull_request --event tag --event deployment --event comment
  3. Add a repository with a longer build timeout.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --timeout 90
  4. Add a repository when config or environment variables are set.
    $ {{.HelpName}} --event push --event pull_request
  5. Add a repository with a starting build number.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --counter 90
  6. Add a repository with a starlark pipeline file.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --pipeline-type starlark
  7. Add a repository with approve build setting set to fork-no-write.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --approve-build fork-no-write

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/repo/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to create a repo.
func add(_ context.Context, c *cli.Command) error {
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

	// create the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config
	r := &repo.Config{
		Action:       internal.ActionAdd,
		Org:          c.String(internal.FlagOrg),
		Name:         c.String(internal.FlagRepo),
		Branch:       c.String("branch"),
		Link:         c.String("link"),
		Clone:        c.String("clone"),
		Visibility:   c.String("visibility"),
		BuildLimit:   c.Int32("build.limit"),
		Timeout:      c.Int32("timeout"),
		Counter:      c.Int64("counter"),
		Private:      internal.StringToBool(c.String("private")),
		Trusted:      internal.StringToBool(c.String("trusted")),
		Active:       internal.StringToBool(c.String("active")),
		Events:       c.StringSlice("event"),
		PipelineType: c.String("pipeline-type"),
		ApproveBuild: c.String("approve-build"),
		Output:       c.String(internal.FlagOutput),
		Color:        output.ColorOptionsFromCLIContext(c),
	}

	// validate repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Validate
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Add
	return r.Add(client)
}
