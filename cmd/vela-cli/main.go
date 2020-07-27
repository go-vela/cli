// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/version"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	// CLI Information

	app.Name = "vela"
	app.HelpName = "vela"
	app.Usage = "CLI for interacting with Vela and managing resources"
	app.Copyright = "Copyright (c) 2020 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// CLI Metadata

	app.Before = load
	app.Compiled = time.Now()
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.Version = version.Version.String()

	// CLI Commands

	app.Commands = []*cli.Command{
		action.Login,
		addCmds,
		chownCmds,
		generateCmds,
		getCmds,
		removeCmds,
		repairCmds,
		restartCmds,
		updateCmds,
		validateCmds,
		viewCmds,
	}

	// CLI Flags

	app.Flags = []cli.Flag{

		// Config Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_CONFIG", "CONFIG_FILE"},
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to Vela configuration file",
			Value:   fmt.Sprintf("%s/.vela/config.yml", os.Getenv("HOME")),
		},

		// API Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ADDR", "VELA_SERVER", "CONFIG_ADDR", "SERVER_ADDR"},
			Name:    internal.FlagAPIAddress,
			Aliases: []string{"a"},
			Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN", "SERVER_TOKEN"},
			Name:    internal.FlagAPIToken,
			Aliases: []string{"t"},
			Usage:   "token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_API_VERSION", "CONFIG_API_VERSION", "API_VERSION"},
			Name:    internal.FlagAPIVersion,
			Aliases: []string{"av"},
			Usage:   "API version for communication with the Vela server",
			Value:   "v1",
		},

		// Log Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL", "LOG_LEVEL"},
			Name:    internal.FlagLogLevel,
			Aliases: []string{"l"},
			Usage:   "set the level of logging - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:   "info",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "CONFIG_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "set the type of output - options: (json|spew|yaml)",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "CONFIG_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "CONFIG_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the CLI",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "CONFIG_ENGINE", "SECRET_ENGINE"},
			Name:    internal.FlagSecretEngine,
			Aliases: []string{"e"},
			Usage:   "provide the secret engine for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "CONFIG_TYPE", "SECRET_TYPE"},
			Name:    internal.FlagSecretType,
			Aliases: []string{"ty"},
			Usage:   "provide the secret type for the CLI",
		},
	}

	// CLI Start

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
