// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/go-vela/cli/cmd/build"
	"github.com/go-vela/cli/cmd/deployment"
	"github.com/go-vela/cli/cmd/hook"
	"github.com/go-vela/cli/cmd/repo"
	"github.com/go-vela/cli/cmd/secret"
	"github.com/go-vela/cli/cmd/service"
	"github.com/go-vela/cli/cmd/step"

	"github.com/urfave/cli/v2"
)

// getCmds defines the commands for listing resources.
var getCmds = cli.Command{
	Name:        "get",
	Category:    "Resource Management",
	Aliases:     []string{"g"},
	Description: "Use this command to list resources for Vela.",
	Usage:       "List resources for Vela via subcommands",
	Subcommands: []*cli.Command{
		&build.GetCmd,
		&deployment.GetCmd,
		&repo.GetCmd,
		&secret.GetCmd,
		&service.GetCmd,
		&step.GetCmd,
		&hook.GetCmd,
	},
}
