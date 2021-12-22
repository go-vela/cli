// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// nolint: dupl // ignore similar code with add
package repo

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// CommandUpdate defines the command for modifying a repository.
var CommandUpdate = &cli.Command{
	Name:        "repo",
	Description: "Use this command to update a repository.",
	Usage:       "Update a new repository from the provided configuration",
	Action:      update,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "REPO_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the repository",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "REPO_NAME"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the name for the repository",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_BRANCH", "REPO_BRANCH"},
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "default branch for the repository",
			Value:   "master",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_LINK", "REPO_LINK"},
			Name:    "link",
			Aliases: []string{"l"},
			Usage:   "full URL to repository in source control",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_CLONE", "REPO_CLONE"},
			Name:    "clone",
			Aliases: []string{"c"},
			Usage:   "full clone URL to repository in source control",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_VISIBILITY", "REPO_VISIBILITY"},
			Name:    "visibility",
			Aliases: []string{"v"},
			Usage:   "access level required to view the repository",
			Value:   constants.VisibilityPublic,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD_LIMIT", "REPO_BUILD_LIMIT"},
			Name:    "build.limit",
			Usage:   "limit of concurrent builds allowed in repository",
			Value:   constants.BuildLimitDefault,
		},
		&cli.Int64Flag{
			EnvVars: []string{"VELA_TIMEOUT", "REPO_TIMEOUT"},
			Name:    "timeout",
			Aliases: []string{"t"},
			Usage:   "max time allowed per build in repository",
			Value:   constants.BuildTimeoutDefault,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_COUNTER", "REPO_COUNTER"},
			Name:    "counter",
			Aliases: []string{"ct"},
			Usage:   "set a value for a new build number",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_PRIVATE", "REPO_PRIVATE"},
			Name:    "private",
			Aliases: []string{"p"},
			Usage:   "disable public access to the repository",
			Value:   "false",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TRUSTED", "REPO_TRUSTED"},
			Name:    "trusted",
			Aliases: []string{"tr"},
			Usage:   "elevated permissions for builds executed for repo",
			Value:   "false",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ACTIVE", "REPO_ACTIVE"},
			Name:    "active",
			Aliases: []string{"a"},
			Usage:   "current status of the repository",
			Value:   "true",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_EVENTS", "REPO_EVENTS"},
			Name:    "event",
			Aliases: []string{"e"},
			Usage:   "webhook event(s) repository responds to",
			Value: cli.NewStringSlice(
				constants.EventPush,
				constants.EventPull,
			),
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_PIPELINE_TYPE", "PIPELINE_TYPE"},
			Name:    "pipeline-type",
			Aliases: []string{"pt"},
			Usage:   "type of base pipeline for the compiler to render",
			Value:   constants.PipelineTypeYAML,
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "REPO_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	// nolint: lll // ignore long line length due to flags
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update a repository with push and pull request enabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --event push --event pull_request
  2. Update a repository with all event types enabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --event push --event pull_request --event tag --event deployment --event comment
  3. Update a repository with a longer build timeout.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --timeout 90
  4. Update a repository when config or environment variables are set.
    $ {{.HelpName}}
  5. Update a repository with a new build number.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --counter 200

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/repo/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to modify a repository.
func update(c *cli.Context) error {
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
		Action:       internal.ActionUpdate,
		Org:          c.String(internal.FlagOrg),
		Name:         c.String(internal.FlagRepo),
		Branch:       c.String("branch"),
		Link:         c.String("link"),
		Clone:        c.String("clone"),
		Visibility:   c.String("visibility"),
		BuildLimit:   c.Int("build.limit"),
		Timeout:      c.Int64("timeout"),
		Counter:      c.Int("counter"),
		Private:      c.Bool("private"),
		Trusted:      c.Bool("trusted"),
		Active:       c.Bool("active"),
		Events:       c.StringSlice("event"),
		PipelineType: c.String("pipeline-type"),
		Output:       c.String(internal.FlagOutput),
	}

	// validate repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Validate
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the update call for the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Update
	return r.Update(client)
}
