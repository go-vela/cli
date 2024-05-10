// SPDX-License-Identifier: Apache-2.0

package output

import (
	"github.com/go-vela/cli/internal"
	"github.com/urfave/cli/v2"
)

// ColorOptions defines the output color options used for syntax highlighting.
type ColorOptions struct {
	Enabled bool
	Theme   string
	Format  string
}

// ColorOptionsFromCLIContext creates a ColorOptions from a CLI context.
func ColorOptionsFromCLIContext(c *cli.Context) ColorOptions {
	opts := ColorOptions{
		Enabled: true,
		Format:  "terminal256",
		Theme:   "monokai",
	}

	opts.Enabled = c.Bool(internal.FlagColor)

	if c.IsSet(internal.FlagColorFormat) {
		opts.Format = c.String(internal.FlagColorFormat)
	}

	if c.IsSet(internal.FlagColorTheme) {
		opts.Theme = c.String(internal.FlagColorTheme)
	}

	return opts
}
