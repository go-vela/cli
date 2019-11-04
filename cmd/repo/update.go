// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"

	"github.com/go-vela/go-vela/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/urfave/cli"
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
		cli.StringFlag{
			Name:   "org",
			Usage:  "Provide the organization for the repository",
			EnvVar: "REPO_ORG",
		},
		cli.StringFlag{
			Name:   "repo",
			Usage:  "Provide the repository contained with the organization",
			EnvVar: "REPO_NAME",
		},

		// optional flags that can be supplied to a command
		cli.StringFlag{
			Name:   "link",
			Usage:  "Link to repository in source control",
			EnvVar: "REPO_LINK",
		},
		cli.StringFlag{
			Name:   "clone",
			Usage:  "Clone link to repository in source control",
			EnvVar: "REPO_CLONE",
		},
		cli.Int64Flag{
			Name:   "timeout",
			Usage:  "Allow management of timeouts",
			EnvVar: "REPO_TIMEOUT",
			Value:  60,
		},
		cli.BoolFlag{
			Name:   "private",
			Usage:  "Allow management of private repositories",
			EnvVar: "REPO_PRIVATE",
		},
		cli.BoolFlag{
			Name:   "trusted",
			Usage:  "Allow management of trusted repositories",
			EnvVar: "REPO_TRUSTED",
		},
		cli.BoolTFlag{
			Name:   "active",
			Usage:  "Allow management of activity on repositories",
			EnvVar: "REPO_ACTIVE",
		},
		cli.StringSliceFlag{
			Name:   "event",
			Usage:  "Allow management of the repository trigger events",
			EnvVar: "REPO_EVENT",
			Value:  &cli.StringSlice{},
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Update a repository.
    $ {{.HelpName}} --org github --repo octocat
 2. Update a repository with all event types enabled.
    $ {{.HelpName}} --org github --repo octocat --event push --event pull_request --event tag --event deployment
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

	// update a carval client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

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
	}

	repository, _, err := client.Repo.Update(org, repo, request)
	if err != nil {
		return err
	}

	fmt.Printf("repo \"%s\" was updated \n", repository.GetFullName())

	return nil
}
