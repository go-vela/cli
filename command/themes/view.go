// SPDX-License-Identifier: Apache-2.0

package themes

import (
	"context"
	"fmt"
	"slices"

	"github.com/alecthomas/chroma/v2/styles"
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/output"
)

// CommandView defines the command for viewing available color themes.
var CommandView = &cli.Command{
	Name:        "themes",
	Description: "Use this command to view available color themes.",
	Usage:       "View available color themes for syntax highlighting",
	Action:      view,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "THEME_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View available color themes.
    $ {{.FullName}}
  2. View available color themes with JSON output.
    $ {{.FullName}} --output json
  3. View available color themes with YAML output.
    $ {{.FullName}} --output yaml

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/themes/view/
`, cli.CommandHelpTemplate),
}

// helper function to view available color themes.
func view(_ context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// get all available theme names from chroma
	themeNames := styles.Names()

	// sort them alphabetically for consistent output
	slices.Sort(themeNames)

	// handle the output based off the provided configuration
	switch c.String(internal.FlagOutput) {
	case output.DriverDump:
		// output the themes in dump format
		return output.Dump(themeNames)
	case output.DriverJSON:
		// output the themes in JSON format
		return output.JSON(themeNames, output.ColorOptionsFromCLIContext(c))
	case output.DriverSpew:
		// output the themes in spew format
		return output.Spew(themeNames)
	case output.DriverYAML:
		// output the themes in YAML format
		return output.YAML(themeNames, output.ColorOptionsFromCLIContext(c))
	default:
		// output the themes in a simple list format
		fmt.Println("Available color themes:")
		fmt.Println()

		for _, theme := range themeNames {
			fmt.Printf("  - %s\n", theme)
		}

		fmt.Println()
		fmt.Println("Use --color.theme <theme-name> to set a theme")
		fmt.Println("Or set in config file: vela config update --color.theme <theme-name>")

		return nil
	}
}
