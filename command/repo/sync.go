// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandSync defines the command for syncing Vela Database with SCM repositories.
var CommandSync = &cli.Command{
	Name:        "repo",
	Aliases:     []string{"repos"},
	Description: "Use this command to sync a repo with the SCM",
	Usage:       "Clean up mismatches between Vela and SCM",
	Action:      sync,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "REPO_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the build",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "REPO_NAME"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the build",
		},

		// Flag to sync all repos in the org

		&cli.BoolFlag{
			EnvVars: []string{"VELA_SYNC_ALL", "SYNC_ALL"},
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "flag to sync all repos in an org",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
	EXAMPLES:
	1. Sync a single repo with SCM.
	  $ {{.HelpName}} --org MyOrg --repo MyRepo
	2. Sync every repo within an org
	  $ {{.HelpName}} --org MyOrg --all
  
    DOCUMENTATION:
  
	https://go-vela.github.io/docs/reference/cli/repo/sync/
  `, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to sync DB with SCM.
func sync(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// parse the Vela client from the context
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/client?tab=doc#Parse
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// check if the all flag has been set
	if c.Bool("all") {
		// create the repo configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config
		r := &repo.Config{
			Action: internal.ActionSyncAll,
			Org:    c.String(internal.FlagOrg),
		}
		// validate repo configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Validate
		err := r.Validate()
		if err != nil {
			return err
		}
		// execute the get call for the repo configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.SyncAll
		return r.SyncAll(client)
	}

	// create the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config
	r := &repo.Config{
		Action: internal.ActionSync,
		Org:    c.String(internal.FlagOrg),
		Name:   c.String(internal.FlagRepo),
	}

	// validate repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Validate
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Sync
	return r.Sync(client)
}
