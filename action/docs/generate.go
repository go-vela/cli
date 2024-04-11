// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/internal/output"
)

// Generate produces documentation for the CLI.
func (c *Config) Generate(a *cli.App) error {
	logrus.Debug("executing generate for docs configuration")

	// generate the docs based off the provided configuration
	switch {
	case c.Markdown:
		logrus.Tracef("creating markdown documentation")

		// generate the documentation from the application configuration
		markdown, err := a.ToMarkdown()
		if err != nil {
			return err
		}

		// output the markdown in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(markdown)
	case c.Man:
		logrus.Tracef("creating man pages documentation")

		// generate the documentation from the application configuration
		man, err := a.ToMan()
		if err != nil {
			return err
		}

		// output the man pages in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(man)
	default:
		// produce an invalid shell error by default
		return fmt.Errorf("invalid documentation format provided")
	}
}
