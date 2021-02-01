// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/docs"

	"github.com/urfave/cli/v2"
)

// DocsGenerate defines the command for producing documentation.
var DocsGenerate = &cli.Command{
	Name:        "docs",
	Description: "Use this command to generate CLI docs.",
	Usage:       "Generate CLI documentation for repository",
	Action:      docsGenerate,
	Hidden:      true,
	Flags: []cli.Flag{

		// Shell Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_MARKDOWN", "DOCS_MARKDOWN"},
			Name:    "markdown",
			Aliases: []string{"m"},
			Usage:   "generate markdown docs",
			Value:   "false",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_MAN", "DOCS_MAN"},
			Name:    "man",
			Aliases: []string{"mn"},
			Usage:   "generate man page docs",
			Value:   "false",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Generate markdown docs for the CLI.
    $ source <({{.HelpName}} --markdown true)
  2. Generate man page docs for the CLI.
    $ source <({{.HelpName}} --man true)

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/docs/generate/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the cli docs.
func docsGenerate(c *cli.Context) error {
	// load variables from the config file
	err := load(c)
	if err != nil {
		return err
	}

	// create the docs configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/docs?tab=doc#Config
	comp := &docs.Config{
		Action:   generateAction,
		Markdown: c.Bool("markdown"),
		Man:      c.Bool("man"),
	}

	// validate docs configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/docs?tab=doc#Config.Validate
	err = comp.Validate()
	if err != nil {
		return err
	}

	// execute the generate call for the docs configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/docs?tab=doc#Config.Generate
	return comp.Generate(c.App)
}
