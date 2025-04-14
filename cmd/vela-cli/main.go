// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/login"
	_version "github.com/go-vela/cli/command/version"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/version"
)

func main() {
	// CLI Information
	cmd := cli.Command{
		Name:                   "vela",
		Version:                version.New().Semantic(),
		Authors:                []any{"Vela Admins"},
		Copyright:              "Copyright 2019 Target Brands, Inc. All rights reserved.",
		Usage:                  "CLI for interacting with Vela and managing resources",
		Before:                 load,
		EnableShellCompletion:  true,
		UseShortOptionHandling: true,
	}

	// CLI Commands

	cmd.Commands = []*cli.Command{
		login.CommandLogin,
		_version.CommandVersion,
		addCmds,
		approveCmds,
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
		syncCmds,
		updateCmds,
		validateCmds,
		viewCmds,
	}

	// CLI Flags

	cmd.Flags = []cli.Flag{

		// Config Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_CONFIG", "CONFIG_FILE"),
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to Vela configuration file",
			Value:   fmt.Sprintf("%s/.vela/config.yml", os.Getenv("HOME")),
		},

		// API Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ADDR", "VELA_SERVER", "CONFIG_ADDR", "SERVER_ADDR"),
			Name:    internal.FlagAPIAddress,
			Aliases: []string{"a"},
			Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_TOKEN", "CONFIG_TOKEN", "SERVER_TOKEN"),
			Name:    internal.FlagAPIToken,
			Aliases: []string{"t"},
			Usage:   "token used for communication with the Vela server",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ACCESS_TOKEN", "CONFIG_ACCESS_TOKEN", "SERVER_ACCESS_TOKEN"),
			Name:    internal.FlagAPIAccessToken,
			Aliases: []string{"at"},
			Usage:   "access token used for communication with the Vela server",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REFRESH_TOKEN", "CONFIG_REFRESH_TOKEN", "SERVER_REFRESH_TOKEN"),
			Name:    internal.FlagAPIRefreshToken,
			Aliases: []string{"rt"},
			Usage:   "refresh access token used for communication with the Vela server",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_API_VERSION", "CONFIG_API_VERSION", "API_VERSION"),
			Name:    internal.FlagAPIVersion,
			Aliases: []string{"av"},
			Usage:   "API version for communication with the Vela server",
			Value:   "v1",
		},

		// Log Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL", "LOG_LEVEL"),
			Name:    internal.FlagLogLevel,
			Aliases: []string{"l"},
			Usage:   "set the level of logging - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:   "info",
		},

		// No Git Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_NO_GIT", "CONFIG_NO_GIT", "NO_GIT"),
			Name:    internal.FlagNoGit,
			Aliases: []string{"ng"},
			Usage:   "set the status of syncing git repo and org with .git/ directory",
			Value:   "false",
		},

		// Color Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_COLOR"),
			Name:    internal.FlagColor,
			Usage:   "enable or disable color output",
		},

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_COLOR_FORMAT"),
			Name:    internal.FlagColorFormat,
			Usage:   "overrides the output color format (default: terminal256)",
		},

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_COLOR_THEME"),
			Name:    internal.FlagColorTheme,
			Usage:   "configures the output color theme (default: monokai)",
		},
	}

	// CLI Start

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
