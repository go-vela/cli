// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"fmt"

	"github.com/sirupsen/logrus"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/internal/output"
)

// Generate produces documentation for the CLI.
func (c *Config) Generate(cmd *cli.Command) error {
	logrus.Debug("executing generate for docs configuration")

	// generate the docs based off the provided configuration
	switch {
	case c.Markdown:
		logrus.Tracef("creating markdown documentation")

		// generate the documentation from the application configuration
		markdown, err := docs.ToMarkdown(cmd)
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
		man, err := docs.ToMan(cmd)
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
