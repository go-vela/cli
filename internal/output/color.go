// SPDX-License-Identifier: Apache-2.0

package output

import (
	"bytes"

	chroma "github.com/alecthomas/chroma/v2/quick"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/internal"
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

// Highlight uses chroma to highlight the provided yaml string.
func Highlight(str string, lexer string, opts ColorOptions) string {
	if opts.Enabled {
		buf := new(bytes.Buffer)

		err := chroma.Highlight(buf, str, lexer, opts.Format, opts.Theme)
		if err == nil {
			str = buf.String()
		} else {
			logrus.Warnf("unable to highlight output: %v", err)
		}
	}

	return str
}
