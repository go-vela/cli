// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/docs"
	"github.com/go-vela/cli/internal"
)

// CommandGenerate defines the command for producing documentation.
var CommandGenerate = &cli.Command{
	Name:        "docs",
	Description: "Use this command to generate CLI docs.",
	Usage:       "Generate CLI documentation for repository",
	Action:      generate,
	Hidden:      true,
	Flags: []cli.Flag{

		// Shell Flags

		&cli.BoolFlag{
			Sources: cli.EnvVars("VELA_MARKDOWN", "DOCS_MARKDOWN"),
			Name:    "markdown",
			Aliases: []string{"m"},
			Usage:   "generate markdown docs",
			Value:   false,
		},
		&cli.BoolFlag{
			Sources: cli.EnvVars("VELA_MAN", "DOCS_MAN"),
			Name:    "man",
			Aliases: []string{"mn"},
			Usage:   "generate man page docs",
			Value:   false,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Generate markdown docs for the CLI.
    $ source <({{.FullName}} --markdown true)
  2. Generate man page docs for the CLI.
    $ source <({{.FullName}} --man true)

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/docs/generate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the cli docs.
func generate(_ context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// create the docs configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/docs?tab=doc#Config
	d := &docs.Config{
		Action:   internal.ActionGenerate,
		Markdown: c.Bool("markdown"),
		Man:      c.Bool("man"),
	}

	// validate docs configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/docs?tab=doc#Config.Validate
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the generate call for the docs configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/docs?tab=doc#Config.Generate
	return d.Generate(c)
}
