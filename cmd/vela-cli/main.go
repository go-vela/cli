// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-vela/cli/command/login"
	_version "github.com/go-vela/cli/command/version"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/version"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// nolint: funlen // ignore function length due to flags
func main() {
	app := cli.NewApp()

	// CLI Information

	app.Name = "vela"
	app.HelpName = "vela"
	app.Usage = "CLI for interacting with Vela and managing resources"
	app.Copyright = "Copyright (c) 2021 Target Brands, Inc. All rights reserved."
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
	app.Version = version.New().Semantic()

	// CLI Commands

	app.Commands = []*cli.Command{
		login.CommandLogin,
		_version.CommandVersion,
		addCmds,
		cancelCmds,
		chownCmds,
		compileCmds,
		execCmds,
		expandCmds,
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
			EnvVars: []string{"VELA_ACCESS_TOKEN", "CONFIG_ACCESS_TOKEN", "SERVER_ACCESS_TOKEN"},
			Name:    internal.FlagAPIAccessToken,
			Aliases: []string{"at"},
			Usage:   "access token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REFRESH_TOKEN", "CONFIG_REFRESH_TOKEN", "SERVER_REFRESH_TOKEN"},
			Name:    internal.FlagAPIRefreshToken,
			Aliases: []string{"rt"},
			Usage:   "refresh access token used for communication with the Vela server",
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

		// No Git Flags

		&cli.BoolFlag{
			EnvVars: []string{"VELA_NO_GIT", "CONFIG_NO_GIT", "NO_GIT"},
			Name:    internal.FlagNoGit,
			Aliases: []string{"gs"},
			Usage:   "set the status of syncing git repo and org with .git/ directory",
			Value:   false,
		},
	}

	// CLI Start

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
