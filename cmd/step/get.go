// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"

	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

// GetCmd defines the command for getting a list of steps.
var GetCmd = cli.Command{
	Name:        "step",
	Aliases:     []string{"steps"},
	Description: "Use this command to get a list of steps.",
	Usage:       "Display a list of steps",
	Action:      get,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"BUILD_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained with the organization",
			EnvVars: []string{"BUILD_REPO"},
		},
		&cli.IntFlag{
			Name:    "build-number",
			Aliases: []string{"build", "b"},
			Usage:   "Provide the build number",
			EnvVars: []string{"BUILD_NUMBER"},
		},

		// optional flags that can be supplied to a command
		&cli.IntFlag{
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "Print a specific page of steps",
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
 1. Get steps for a build.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --build-number 1
 2. Get steps for a build with wide view output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --build-number 1 --output wide
 3. Get steps for a build with yaml output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --build-number 1 --output yaml
 4. Get steps for a build with json output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --build-number 1 --output json
 5. Get steps for a build when org and repo config or environment variables are set.
    $ {{.HelpName}} --number 1
`, cli.CommandHelpTemplate),
}

// helper function to execute vela get steps cli command
func get(c *cli.Context) error {
	// get org, repo and number information from cmd flags
	org, repo, number := c.String("org"), c.String("repo"), c.Int("build-number")

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

	steps, _, err := client.Step.GetAll(org, repo, number, opts)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(steps, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "yaml":
		output, err := yaml.Marshal(steps)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "wide":
		table := uitable.New()
		table.MaxColWidth = 200
		table.Wrap = true
		// spaces after status widen column for better readability
		table.AddRow("NUMBER", "NAME", "STATUS", "CREATED", "FINISHED", "DURATION")

		for _, s := range reverse(*steps) {
			if s.GetStatus() == constants.StatusRunning {
				table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), s.GetExitCode(), s.GetRuntime(), humanize.Time(time.Unix(s.GetCreated(), 0)), humanize.Time(time.Unix(s.GetFinished(), 0)), "...")
			} else {
				table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), s.GetExitCode(), s.GetRuntime(), humanize.Time(time.Unix(s.GetCreated(), 0)), humanize.Time(time.Unix(s.GetFinished(), 0)), calcDuration(&s))
			}
		}

		fmt.Println(table)

	default:
		table := uitable.New()
		table.MaxColWidth = 50
		table.Wrap = true
		// spaces after status widen column for better readability
		table.AddRow("NUMBER", "NAME", "STATUS", "DURATION")

		for _, s := range reverse(*steps) {
			if s.GetStatus() == constants.StatusRunning {
				table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), "...")
			} else {
				table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), calcDuration(&s))
			}
		}

		fmt.Println(table)
	}

	return nil
}

// calcDuration gets build duration
func calcDuration(b *library.Step) string {
	dur := (b.GetFinished() - b.GetStarted())

	if dur < 60 {
		return fmt.Sprintf("%ds", dur)
	}

	return fmt.Sprintf("%dm", dur/60)
}

// helper function to reverse the build list output
func reverse(s []library.Step) []library.Step {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].GetNumber() < s[j].GetNumber()
	})

	return s
}
