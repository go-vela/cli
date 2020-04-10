// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/urfave/cli/v2"
)

// UpdateCmd defines the command to update a repository.
var UpdateCmd = cli.Command{
	Name:        "repo",
	Description: "Use this command to update a repository.",
	Usage:       "Update a repository",
	Action:      update,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"REPO_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained with the organization",
			EnvVars: []string{"REPO_NAME"},
		},

		// optional flags that can be supplied to a command
		&cli.StringFlag{
			Name:    "link",
			Usage:   "Link to repository in source control",
			EnvVars: []string{"REPO_LINK"},
		},
		&cli.StringFlag{
			Name:    "clone",
			Usage:   "Clone link to repository in source control",
			EnvVars: []string{"REPO_CLONE"},
		},
		&cli.Int64Flag{
			Name:    "timeout",
			Usage:   "Allow management of timeouts",
			EnvVars: []string{"REPO_TIMEOUT"},
			Value:   60,
		},
		&cli.BoolFlag{
			Name:    "private",
			Usage:   "Allow management of private repositories",
			EnvVars: []string{"REPO_PRIVATE"},
		},
		&cli.BoolFlag{
			Name:    "trusted",
			Usage:   "Allow management of trusted repositories",
			EnvVars: []string{"REPO_TRUSTED"},
		},
		&cli.BoolFlag{
			Name:    "active",
			Usage:   "Allow management of activity on repositories",
			EnvVars: []string{"REPO_ACTIVE"},
			Value:   true,
		},
		&cli.StringSliceFlag{
			Name:    "event",
			Usage:   "Allow management of the repository trigger events",
			EnvVars: []string{"REPO_EVENT"},
			Value:   &cli.StringSlice{},
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Update a repository.
    $ {{.HelpName}} --org github --repo octocat
 2. Update a repository with all event types enabled.
    $ {{.HelpName}} --org github --repo octocat --event push --event pull_request --event tag --event deployment --event comment
 3. Update a repository with a longer build timeout.
    $ {{.HelpName}} --org github --repo octocat --timeout 90
 4. Update a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to execute a update repo cli command
func update(c *cli.Context) error {
	// get org and repo information from cmd flags
	org, repo := c.String("org"), c.String("repo")

	// update a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// resource to update on server
	request := &library.Repo{
		FullName: vela.String(fmt.Sprintf("%s/%s", org, repo)),
		Org:      vela.String(org),
		Name:     vela.String(repo),
		Link:     vela.String(c.String("link")),
		Clone:    vela.String(c.String("clone")),
		Timeout:  vela.Int64(c.Int64("timeout")),
		Private:  vela.Bool(c.Bool("private")),
		Trusted:  vela.Bool(c.Bool("trusted")),
		Active:   vela.Bool(c.Bool("active")),
	}

	for _, event := range c.StringSlice("event") {
		if event == constants.EventPush {
			request.AllowPush = vela.Bool(true)
		}

		if event == constants.EventPull {
			request.AllowPull = vela.Bool(true)
		}

		if event == constants.EventTag {
			request.AllowTag = vela.Bool(true)
		}

		if event == constants.EventDeploy {
			request.AllowDeploy = vela.Bool(true)
		}

		if event == constants.EventComment {
			request.AllowComment = vela.Bool(true)
		}
	}

	repository, _, err := client.Repo.Update(org, repo, request)
	if err != nil {
		return err
	}

	fmt.Printf("repo \"%s\" was updated \n", repository.GetFullName())

	return nil
}
