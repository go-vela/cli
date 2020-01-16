// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-vela/cli/cmd"
	"github.com/go-vela/cli/version"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
)

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version,v,V",
		Usage: "print the CLI version information",
	}

	app := cli.NewApp()
	app.Name = "vela"
	app.HelpName = "vela"
	app.Usage = "CLI for managing Vela resources"
	app.Compiled = time.Now()
	app.Version = version.Version.String()
	app.Copyright = "Copyright (c) 2020 Target Brands, Inc. All rights reserved."
	app.Authors = []cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	app.Commands = cmd.Vela
	app.Flags = []cli.Flag{
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_ADDR",
			Name:   "addr",
			Usage:  "location of vela server",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_TOKEN",
			Name:   "token",
			Usage:  "User token for Vela server",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_API_VERSION",
			Name:   "api-version",
			Usage:  "api version to use for Vela server",
			Value:  "v1",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_LOG_LEVEL",
			Name:   "log-level",
			Usage:  "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:  "info",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_BUILD_ORG",
			Name:   "org",
			Usage:  "Provide the organization owner for the repository",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_BUILD_REPO",
			Name:   "repo",
			Usage:  "Provide the repository contained within the organization",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_SECRET_ENGINE",
			Name:   "secret-engine",
			Usage:  "Provide the engine for where the secret to be stored",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			EnvVar: "VELA_SECRET_TYPE",
			Name:   "secret-type",
			Usage:  "Provide the kind of secret to be stored",
		}),
		cli.StringFlag{
			EnvVar: "VELA_CONFIG",
			Name:   "config",
			Usage:  "path to Vela configuration file",
			Value:  fmt.Sprintf("%s/.vela/config.yml", os.Getenv("HOME")),
		},
	}
	app.Before = load

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

// load is a helper function that loads the necessary configuration for the CLI.
func load(c *cli.Context) error {
	config := c.GlobalString("config")

	_, err := os.Stat(config)
	if err == nil {
		err = altsrc.InitInputSourceWithContext(c.App.Flags, func(context *cli.Context) (altsrc.InputSourceContext, error) {
			yaml, err := altsrc.NewYamlSourceFromFile(config)
			return yaml, err
		})(c)
		if err != nil {
			return fmt.Errorf("Unable to load config file @ %s: %v", config, err)
		}
	}
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("Unable to search for config file @ %s: %v", config, err)
		}
		logrus.Warningf("Unable to find config file @ %s", config)
	}

	err = validate(c)
	if err != nil {
		return err
	}

	switch c.GlobalString("log-level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "i", "info", "Info", "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	}

	return nil
}

// validate is a helper function that ensures the necessary configuration is set for the CLI.
func validate(c *cli.Context) error {
	args := c.Args()
	// DO NOT validate if help argument is provided
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return nil
		}
	}

	// DO NOT validate if config argument is provided
	if args.Get(1) == "config" || args.Get(1) == "c" {
		return nil
	}

	// DO NOT validate if login argument is provided
	if args.Get(0) == "login" || args.Get(0) == "l" {
		return nil
	}

	if len(c.GlobalString("addr")) == 0 {
		return fmt.Errorf("No vela address provided")
	}

	if len(c.GlobalString("token")) == 0 {
		return fmt.Errorf("No vela token provided")
	}

	return nil
}
