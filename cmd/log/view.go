// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"fmt"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// ViewCmd defines the command for viewing the logs from a build or step.
var ViewCmd = cli.Command{
	Name:        "log",
	Aliases:     []string{"logs"},
	Description: "Use this command to capture the logs from a build or step.",
	Usage:       "View logs from the provided build or step",
	Action:      view,
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
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   "Provide the log type service/step",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. View steps logs of a build for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --type step
      OR
    $ {{.HelpName}} --build-number 1 --type step      (when org and repo config environment variables are set).

 2. View services logs of a build for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --type service
      OR
    $ {{.HelpName}} --build-number 1 --type service   (when org and repo config environment variables are set).
`, cli.CommandHelpTemplate),
}

// helper function to execute logs cli command
func view(c *cli.Context) error {

	// get org, repo and number information from cmd flags
	org, repo, number, logType := c.String("org"), c.String("repo"), c.Int("build-number"), c.String("type")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// Get the build you just created
	build, _, err := client.Build.Get(org, repo, number)
	if err != nil {
		return err
	}

	switch logType {
	case "service":
		err = PrintServicesLogs(client, org, repo, build.Number)
		if err != nil {
			return err
		}
	case "step":
		err = PrintStepsLogs(client, org, repo, build.Number)
		if err != nil {
			return err
		}
	}

	return nil
}

func PrintServicesLogs(client *vela.Client, org string, repo string, buildNumber *int) error {

	// Get all services for the build
	services, _, err := client.Svc.GetAll(org, repo, *buildNumber, nil)
	if err != nil {
		return err
	}

	// Print logs for each service
	for _, service := range *services {
		serviceLog, _, err := client.Log.GetService(org, repo, *buildNumber, *service.Number)
		if err != nil {
			return err
		}
		fmt.Printf("%s \n", serviceLog.GetData())
	}

	return nil
}

func PrintStepsLogs(client *vela.Client, org string, repo string, buildNumber *int) error {

	// Get all steps for the build
	steps, _, err := client.Step.GetAll(org, repo, *buildNumber, nil)
	if err != nil {
		return err
	}

	// Print logs for each steps
	for _, step := range *steps {
		stepLog, _, err := client.Log.GetStep(org, repo, *buildNumber, *step.Number)
		if err != nil {
			return err
		}
		fmt.Printf("%s \n", stepLog.GetData())
	}

	return nil
}
