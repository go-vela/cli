package hook

import (
	"fmt"
	"github.com/go-vela/sdk-go/vela"
	"github.com/urfave/cli/v2"
)

var GetCmd = cli.Command{

	Name:        "hook",
	Aliases:     []string{"hooks"},
	Description: "Use this command to get all hooks for a repo",
	Usage:       "Display list of hooks",
	Action:      get,
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
 1. Get all hooks for a repository.
    $ {{.HelpName}} --org github --repo octocat 
	OR
	$ {{.HelpName}} (When org & repo env variables are set)

 2. Get all hooks for a repository with different output formats.
    $ {{.HelpName}} --org github --repo octocat --output json
    $ {{.HelpName}} --org github --repo octocat --output yaml
    $ {{.HelpName}} --org github --repo octocat --output wide
`, cli.CommandHelpTemplate),
}

func get(c *cli.Context) error {
	org, repo, format := c.String("org"), c.String("repo"), c.String("output")

	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	client.Authentication.SetTokenAuth(c.String("token"))

	hooks, _, err := client.Hook.GetAll(org, repo, nil)
	if err != nil {
		return err
	}

	err = PrintOutput(format, *hooks...)
	if err != nil {
		return err
	}

	return nil
}
