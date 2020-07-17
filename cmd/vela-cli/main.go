// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal/client"
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
			Name:    client.KeyAddress,
			Aliases: []string{"a"},
			Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TOKEN", "CONFIG_TOKEN", "SERVER_TOKEN"},
			Name:    client.KeyToken,
			Aliases: []string{"t"},
			Usage:   "token used for communication with the Vela server",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_API_VERSION", "CONFIG_API_VERSION", "API_VERSION"},
			Name:    "api.version",
			Aliases: []string{"av"},
			Usage:   "API version for communication with the Vela server",
			Value:   "v1",
		},

		// Log Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "CONFIG_LOG_LEVEL", "LOG_LEVEL"},
			Name:    "log.level",
			Aliases: []string{"l"},
			Usage:   "set the level of logging - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:   "info",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "CONFIG_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "set the type of output - options: (json|spew|yaml)",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "CONFIG_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "CONFIG_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the CLI",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "CONFIG_ENGINE", "SECRET_ENGINE"},
			Name:    "secret.engine",
			Aliases: []string{"e"},
			Usage:   "provide the secret engine for the CLI",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "CONFIG_TYPE", "SECRET_TYPE"},
			Name:    "secret.type",
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

// load is a helper function that loads the necessary configuration for the CLI.
func load(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action: "load",
		File:   c.String("config"),
	}

	// validate config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Validate
	err := conf.Validate()
	if err == nil {
		// execute the load call for the config file configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config.Load
		err = conf.Load(c)
		if err != nil {
			return err
		}
	}

	// set log level for the CLI
	switch c.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	return nil
}
