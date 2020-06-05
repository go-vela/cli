// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/library"

	"github.com/gosuri/uitable"

	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

// GetCmd defines the command for getting a list of deployments.
var GetCmd = cli.Command{
	Name:        "deployment",
	Aliases:     []string{"deployments"},
	Description: "Use this command to get a list of deployments.",
	Usage:       "Display a list of deployments",
	Action:      get,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"VELA_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained within the organization",
			EnvVars: []string{"VELA_REPO"},
		},

		// optional flags that can be supplied to a command
		&cli.IntFlag{
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "Print a specific page of deployments",
			Value:   1,
		},
		&cli.IntFlag{
			Name:    "per-page",
			Aliases: []string{"pp"},
			Usage:   "Expand the number of items contained within page",
			Value:   10,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in wide, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Get deployments for a repository.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld
 2. Get deployments for a repository with wide view output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --output wide
 3. Get deployments for a repository with yaml output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --output yaml
 4. Get deployments for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --output json
 5. Get deployments for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to execute vela get deployments cli command
func get(c *cli.Context) error {
	// get org and repo information from cmd flags
	org, repo := c.String("org"), c.String("repo")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}
	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// set the page options based on user input
	opts := &vela.ListOptions{
		Page:    c.Int("page"),
		PerPage: c.Int("per-page"),
	}

	deployments, _, err := client.Deployment.GetAll(org, repo, opts)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(deployments, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "yaml":
		output, err := yaml.Marshal(deployments)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "wide":
		table := uitable.New()
		table.MaxColWidth = 200
		table.Wrap = true
		// spaces after status widen column for better readability
		table.AddRow("ID", "TASK", "USER", "REF", "TARGET", "COMMIT", "URL", "DESCRIPTION")

		for _, d := range reverse(*deployments) {
			table.AddRow(d.GetID(), d.GetTask(), d.GetUser(), d.GetRef(), d.GetTarget(), d.GetCommit(), d.GetURL(), d.GetDescription())
		}

		fmt.Println(table)

	default:
		table := uitable.New()
		table.MaxColWidth = 50
		table.Wrap = true

		table.AddRow("ID", "TASK", "USER", "REF", "TARGET")

		for _, d := range reverse(*deployments) {
			table.AddRow(d.GetID(), d.GetTask(), d.GetUser(), d.GetRef(), d.GetTarget())
		}

		fmt.Println(table)
	}

	return nil
}

// helper function to reverse the deployment list output
func reverse(d []library.Deployment) []library.Deployment {
	sort.SliceStable(d, func(i, j int) bool {
		return d[i].GetID() < d[j].GetID()
	})

	return d
}
