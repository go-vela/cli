// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

	"github.com/urfave/cli/v2"
)

// viewCmds defines the commands for inspecting resources.
var viewCmds = &cli.Command{
	Name:                   "view",
	Category:               "Resource Management",
	Aliases:                []string{"v"},
	Description:            "Use this command to view a resource for Vela.",
	Usage:                  "View details for a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for inspecting a build
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#BuildView
		action.BuildView,

		// add the sub command for inspecting a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#ConfigView
		action.ConfigView,

		// add the sub command for inspecting a deployment
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#DeploymentView
		action.DeploymentView,

		// add the sub command for inspecting a hook
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#HookView
		action.HookView,

		// add the sub command for inspecting a log
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#LogView
		action.LogView,

		// add the sub command for inspecting a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#PipelineView
		action.PipelineView,

		// add the sub command for inspecting a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#RepoView
		action.RepoView,

		// add the sub command for inspecting a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#SecretView
		action.SecretView,

		// add the sub command for inspecting a service
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#ServiceView
		action.ServiceView,

		// add the sub command for inspecting a step
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#StepView
		action.StepView,
	},
}
