// SPDX-License-Identifier: Apache-2.0

package output

import (
	"bytes"
	"os"

	chroma "github.com/alecthomas/chroma/v2/quick"
	"github.com/muesli/termenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"
	"golang.org/x/term"

	"github.com/go-vela/cli/internal"
)

// ColorOptions defines the output color options used for syntax highlighting.
type ColorOptions struct {
	Enabled       bool
	Format        string
	Theme         string
	ThemeLight    string
	UserSpecified bool
}

// GetTheme returns the appropriate theme based on terminal background
// and user preferences.
func (opts ColorOptions) GetTheme() string {
	if opts.UserSpecified {
		logrus.Debug("using user specified color theme")

		return opts.Theme
	}

	if termenv.HasDarkBackground() {
		logrus.Debug("detected dark terminal background, using default theme")

		return opts.Theme
	}

	logrus.Debug("detected light terminal background, using default light theme")

	return opts.ThemeLight
}

// shouldEnableColor determines if color should be enabled based on
// environment variables and CLI flags following standard conventions.
func shouldEnableColor(c *cli.Command) bool {
	// 1. NO_COLOR - if set to any value, disable colors (highest priority)
	//    See: https://no-color.org/
	if _, exists := os.LookupEnv("NO_COLOR"); exists {
		logrus.Debug("NO_COLOR set, colors will be suppressed")

		return false
	}

	// 2. User-specified --color flag takes precedence over env vars
	if c.IsSet(internal.FlagColor) {
		logrus.Debug("--color is set, using supplied value")

		return internal.StringToBool(c.String(internal.FlagColor))
	}

	// 3. CLICOLOR_FORCE - if non-zero, force colors even if not a TTY
	//    See: https://bixense.com/clicolors/
	if cliColorForce := os.Getenv("CLICOLOR_FORCE"); cliColorForce != "" && cliColorForce != "0" {
		logrus.Debug("CLICOLOR_FORCE is set, forcing colors")

		return true
	}

	// 4. CLICOLOR=0 explicitly disables colors
	if cliColor := os.Getenv("CLICOLOR"); cliColor == "0" {
		logrus.Debug("CLICOLOR set to '0', colors will be suppressed")

		return false
	}

	// 5. If not a terminal, don't use color by default
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		logrus.Debug("no TTY, colors will be suppressed")

		return false
	}

	// Default: enable colors
	return true
}

// ColorOptionsFromCLIContext creates a ColorOptions from a CLI context.
func ColorOptionsFromCLIContext(c *cli.Command) ColorOptions {
	opts := ColorOptions{
		Enabled:       shouldEnableColor(c),
		Format:        "terminal256",
		Theme:         "monokai",
		ThemeLight:    "monokailight",
		UserSpecified: false,
	}

	if c.IsSet(internal.FlagColorFormat) {
		logrus.Debug("using supplied color format value")

		opts.Format = c.String(internal.FlagColorFormat)
	}

	if c.IsSet(internal.FlagColorTheme) {
		logrus.Debug("using user supplied color theme")

		opts.Theme = c.String(internal.FlagColorTheme)
		opts.UserSpecified = true
	}

	return opts
}

// Highlight uses chroma to highlight the provided yaml string.
func Highlight(str string, lexer string, opts ColorOptions) string {
	if opts.Enabled {
		buf := new(bytes.Buffer)

		theme := opts.GetTheme()

		err := chroma.Highlight(buf, str, lexer, opts.Format, theme)
		if err == nil {
			str = buf.String()
		} else {
			logrus.Warnf("unable to highlight output: %v", err)
		}
	}

	return str
}
