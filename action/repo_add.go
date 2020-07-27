// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// RepoAdd defines the command for creating a repository.
var RepoAdd = &cli.Command{
	Name:        "repo",
	Description: "Use this command to add a repository.",
	Usage:       "Add a new repository from the provided configuration",
	Action:      repoAdd,
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
			Value:   "public",
		},
		&cli.Int64Flag{
			EnvVars: []string{"VELA_TIMEOUT", "REPO_TIMEOUT"},
			Name:    "timeout",
			Aliases: []string{"t"},
			Usage:   "max time allowed per build in repository",
			Value:   30,
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_PRIVATE", "REPO_PRIVATE"},
			Name:    "private",
			Aliases: []string{"p"},
			Usage:   "disable public access to the repository",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_TRUSTED", "REPO_TRUSTED"},
			Name:    "trusted",
			Aliases: []string{"tr"},
			Usage:   "elevated permissions for builds executed for repo",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_ACTIVE", "REPO_ACTIVE"},
			Name:    "active",
			Aliases: []string{"a"},
			Usage:   "current status of the repository",
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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "REPO_OUTPUT"},
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

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/repo/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// create a repo.
func repoAdd(c *cli.Context) error {
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
		Action:     addAction,
		Org:        c.String(internal.FlagOrg),
		Name:       c.String(internal.FlagRepo),
		Branch:     c.String("branch"),
		Link:       c.String("link"),
		Clone:      c.String("clone"),
		Visibility: c.String("visibility"),
		Timeout:    c.Int64("timeout"),
		Private:    c.Bool("private"),
		Trusted:    c.Bool("trusted"),
		Active:     c.Bool("active"),
		Events:     c.StringSlice("event"),
		Output:     c.String(internal.FlagOutput),
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
