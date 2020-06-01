package hook

import (
	"fmt"
	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"
	"github.com/urfave/cli/v2"
)

var ViewCmd = cli.Command{

	Name:        "hook",
	Description: "Use this command to view a hook",
	Usage:       "View details of the provided hook",
	Action:      view,
	Before:      validate,
	Flags: []cli.Flag{
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
			Name:    "hook-number",
			Aliases: []string{"hook"},
			Usage:   "Provide the hook number",
		},
		// Supports printing the output in json/yaml/wide format.
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			EnvVars: []string{"VELA_OUTPUT"},
			Usage:   "Print the output in json/yaml/wide format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. View a hook of a repository.
    $ {{.HelpName}} --org github --repo octocat --hook-number 1
	OR
	$ {{.HelpName}} --hook-number 1     (When org & repo env variables are set)

 2. View a hook of a repository with different output formats.
    $ {{.HelpName}} --org github --repo octocat --hook-number 1 --output json
    $ {{.HelpName}} --org github --repo octocat --hook-number 1 --output yaml
    $ {{.HelpName}} --org github --repo octocat --hook-number 1 --output wide
`, cli.CommandHelpTemplate),
}

func view(c *cli.Context) error {
	org, repo, hookNumber, outputFormat := c.String("org"), c.String("repo"), c.Int("hook-number"), c.String("output")

	if hookNumber == 0 {
		return util.InvalidCommand("hook-number")
	}
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	client.Authentication.SetTokenAuth(c.String("token"))

	hook, _, err := client.Hook.Get(org, repo, hookNumber)
	if err != nil {
		return err
	}

	err = PrintOutput(outputFormat, *hook)
	if err != nil {
		return err
	}

	return nil
}
