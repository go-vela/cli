// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/build"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// BuildGet defines the command for capturing a list of builds.
var BuildGet = &cli.Command{
	Name:        "build",
	Aliases:     []string{"builds"},
	Description: "Use this command to get a list of builds.",
	Usage:       "Display a list of builds",
	Action:      buildGet,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG"},
			Name:    "org",
			Usage:   "Provide the organization for the build",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO"},
			Name:    "repo",
			Usage:   "Provide the repository for the build",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT"},
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in wide, yaml or json format",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE"},
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "Print a specific page of builds",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE"},
			Name:    "per.page",
			Aliases: []string{"pp"},
			Usage:   "Expand the number of items contained within page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Get builds for a repository.
    $ {{.HelpName}} --org github --repo octocat
 2. Get builds for a repository with wide view output.
    $ {{.HelpName}} --org github --repo octocat --output wide
 3. Get builds for a repository with yaml output.
    $ {{.HelpName}} --org github --repo octocat --output yaml
 4. Get builds for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --output json
 5. Get builds for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of builds.
func buildGet(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the build configuration
	b := &build.Config{
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate build configuration
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the build configuration
	return b.Get(client)
}
