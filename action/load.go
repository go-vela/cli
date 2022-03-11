// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"github.com/go-vela/cli/action/config"
	"github.com/go-vela/cli/internal"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// Load imports the configuration file for the CLI.
func Load(c *cli.Context) error {
	// create the config file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/config?tab=doc#Config
	conf := &config.Config{
		Action: internal.ActionLoad,
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

	// set org and repo
	internal.SetGitConfigContext(c)

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

	return nil
}
