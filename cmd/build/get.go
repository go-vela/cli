// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// GetCmd defines the command for getting a list of builds.
var GetCmd = cli.Command{
	Name:        "build",
	Aliases:     []string{"builds"},
	Description: "Use this command to get a list of builds.",
	Usage:       "Display a list of builds",
	Action:      get,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		cli.StringFlag{
			Name:   "org",
			Usage:  "Provide the organization for the repository",
			EnvVar: "BUILD_ORG",
		},
		cli.StringFlag{
			Name:   "repo",
			Usage:  "Provide the repository contained within the organization",
			EnvVar: "BUILD_REPO",
		},

		// optional flags that can be supplied to a command
		cli.IntFlag{
			Name:  "page,p",
			Usage: "Print a specific page of builds",
			Value: 1,
		},
		cli.IntFlag{
			Name:  "per-page,pp",
			Usage: "Expand the number of items contained within page",
			Value: 10,
		},
		cli.StringFlag{
			Name:  "output,o",
			Usage: "Print the output in wide, yaml or json format",
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

// helper function to execute vela get builds cli command
func get(c *cli.Context) error {
	// get org and repo information from cmd flags
	org, repo := c.String("org"), c.String("repo")

	// create a vela client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}
	// set token from global config
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

	// set the page options based on user input
	opts := &vela.ListOptions{
		Page:    c.Int("page"),
		PerPage: c.Int("per-page"),
	}

	builds, _, err := client.Build.GetAll(org, repo, opts)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(builds, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "yaml":
		output, err := yaml.Marshal(builds)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "wide":
		table := uitable.New()
		table.MaxColWidth = 200
		table.Wrap = true
		// spaces after status widen column for better readability
		table.AddRow("NUMBER", "STATUS", "EVENT", "BRANCH", "COMMIT", "DURATION", "CREATED", "FINISHED", "AUTHOR")

		for _, b := range reverse(*builds) {
			modifyBuild(&b)

			if b.GetStatus() == constants.StatusPending || b.GetStatus() == constants.StatusRunning {
				table.AddRow(b.GetNumber(), b.GetStatus(), b.GetEvent(), b.GetBranch(), b.GetCommit(), "...", humanize.Time(time.Unix(b.GetCreated(), 0)), humanize.Time(time.Unix(b.GetFinished(), 0)), b.GetAuthor())
			} else {
				table.AddRow(b.GetNumber(), b.GetStatus(), b.GetEvent(), b.GetBranch(), b.GetCommit(), calcDuration(&b), humanize.Time(time.Unix(b.GetCreated(), 0)), humanize.Time(time.Unix(b.GetFinished(), 0)), b.GetAuthor())
			}
		}

		fmt.Println(table)

	default:
		table := uitable.New()
		table.MaxColWidth = 50
		table.Wrap = true
		// spaces after status widen column for better readability
		table.AddRow("NUMBER", "STATUS", "EVENT", "BRANCH", "DURATION")

		for _, b := range reverse(*builds) {
			modifyBuild(&b)

			if b.GetStatus() == constants.StatusPending || b.GetStatus() == constants.StatusRunning {
				table.AddRow(b.GetNumber(), b.GetStatus(), b.GetEvent(), b.GetBranch(), "...")
			} else {
				table.AddRow(b.GetNumber(), b.GetStatus(), b.GetEvent(), b.GetBranch(), calcDuration(&b))
			}
		}

		fmt.Println(table)
	}

	return nil
}

// calcDuration gets build duration
func calcDuration(b *library.Build) string {
	dur := (b.GetFinished() - b.GetStarted())

	if dur < 60 {
		return fmt.Sprintf("%ds", dur)
	}

	return fmt.Sprintf("%dm", dur/60)
}

// modifybuild changes the event data within the struct
// to reflect to custom output when we are not outputing yaml or json
func modifyBuild(b *library.Build) {
	switch d := strings.Split(b.GetRef(), "/")[1]; d {
	case "tags":
		*b.Ref = "tag"
		*b.Event = d
	case "pull":
		*b.Ref = "pull_request"
		*b.Branch = *b.Ref
	default:
		*b.Ref = ""
	}
}

// helper function to reverse the build list output
func reverse(b []library.Build) []library.Build {
	sort.SliceStable(b, func(i, j int) bool {
		return b[i].GetNumber() < b[j].GetNumber()
	})

	return b
}
