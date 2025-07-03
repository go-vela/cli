// SPDX-License-Identifier: Apache-2.0

package completion

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/completion"
	"github.com/go-vela/cli/internal"
)

// CommandGenerate defines the command for producing an auto completion script.
var CommandGenerate = &cli.Command{
	Name:        "completion",
	Description: "Use this command to generate a shell auto completion script.",
	Usage:       "Generate a shell auto completion script",
	Action:      generate,
	Flags: []cli.Flag{

		// Shell Flags

		&cli.BoolFlag{
			Sources: cli.EnvVars("VELA_BASH", "COMPLETION_BASH"),
			Name:    "bash",
			Aliases: []string{"b"},
			Usage:   "generate a bash auto completion script",
			Value:   false,
		},
		&cli.BoolFlag{
			Sources: cli.EnvVars("VELA_ZSH", "COMPLETION_ZSH"),
			Name:    "zsh",
			Aliases: []string{"z"},
			Usage:   "generate a zsh auto completion script",
			Value:   false,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Enable auto completion for the current bash session.
    $ source <({{.FullName}} --bash true)
  2. Enable auto completion for the current zsh session.
    $ source <({{.FullName}} --zsh true)
  3. Enable auto completion for bash permanently.
    visit https://go-vela.github.io/docs/reference/cli/completion/generate/#bash
  4. Enable auto completion for zsh permanently.
    visit https://go-vela.github.io/docs/reference/cli/completion/generate/#zsh

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/completion/generate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to produce the
// config file.
func generate(_ context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// create the completion configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/completion?tab=doc#Config
	comp := &completion.Config{
		Action: internal.ActionGenerate,
		Bash:   c.Bool("bash"),
		Zsh:    c.Bool("zsh"),
	}

	// validate completion configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/completion?tab=doc#Config.Validate
	err = comp.Validate()
	if err != nil {
		return err
	}

	// execute the generate call for the completion configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/completion?tab=doc#Config.Generate
	return comp.Generate()
}
