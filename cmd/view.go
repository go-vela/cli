// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/go-vela/cli/cmd/build"
	"github.com/go-vela/cli/cmd/config"
	"github.com/go-vela/cli/cmd/deployment"
	"github.com/go-vela/cli/cmd/hook"
	"github.com/go-vela/cli/cmd/log"
	"github.com/go-vela/cli/cmd/repo"
	"github.com/go-vela/cli/cmd/secret"
	"github.com/go-vela/cli/cmd/service"
	"github.com/go-vela/cli/cmd/step"

	"github.com/urfave/cli/v2"
)

// viewCmds defines the commands for viewing resources.
var viewCmds = cli.Command{
	Name:        "view",
	Category:    "Resource Management",
	Aliases:     []string{"v"},
	Description: "Use this command to view resources for Vela.",
	Usage:       "View resources for Vela via subcommands",
	Subcommands: []*cli.Command{
		&build.ViewCmd,
		&deployment.ViewCmd,
		&config.ViewCmd,
		&log.ViewCmd,
		&repo.ViewCmd,
		&secret.ViewCmd,
		&service.ViewCmd,
		&step.ViewCmd,
		&hook.ViewCmd,
	},
}
