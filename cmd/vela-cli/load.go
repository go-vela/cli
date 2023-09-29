// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// load is a helper function that loads the necessary configuration for the CLI.
func load(c *cli.Context) error {
	// set log level for the CLI
	switch c.String(internal.FlagLogLevel) {
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

	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action: "load",
		File:   c.String(internal.FlagConfig),
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

	return nil
}
