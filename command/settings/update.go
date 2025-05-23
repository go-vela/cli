// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/settings"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/util"
)

const (
	CompilerCloneImageKey        = "compiler.clone-image"
	CompilerTemplateDepthKey     = "compiler.template-depth"
	CompilerStarlarkExecLimitKey = "compiler.starlark-exec-limit"
	QueueRouteAddKey             = "queue.add-route"
	QueueRouteDropKey            = "queue.drop-route"
	SCMRepoRolesMapKey           = "scm.repo.roles-map"
	SCMOrgRolesMapKey            = "scm.org.roles-map"
	SCMTeamRolesMapKey           = "scm.team.roles-map"
	RepoAllowlistAddKey          = "add-repo"
	RepoAllowlistDropKey         = "drop-repo"
	ScheduleAllowAddlistKey      = "add-schedule"
	ScheduleAllowDroplistKey     = "drop-schedule"
	MaxDashboardReposKey         = "max-dashboard-repos"
)

// CommandUpdate defines the command for modifying a settings.
var CommandUpdate = &cli.Command{
	Name:        "settings",
	Description: "(Platform Admin Only) Use this command to update settings.",
	Usage:       "Update settings from the provided configuration",
	Action:      update,
	Flags: []cli.Flag{
		// Queue Flags

		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_QUEUE_ADD_ROUTES", "QUEUE_ADD_ROUTES"),
			Name:    QueueRouteAddKey,
			Aliases: []string{"queue-route", "add-route", "routes", "route", "r"},
			Usage:   "list of routes to add to the queue",
		},

		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_QUEUE_DROP_ROUTES", "QUEUE_DROP_ROUTES"),
			Name:    QueueRouteDropKey,
			Aliases: []string{"drop-route"},
			Usage:   "list of routes to drop from the queue",
		},

		// Compiler Flags

		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_COMPILER_TEMPLATE_DEPTH", "VELA_TEMPLATE_DEPTH", "TEMPLATE_DEPTH"),
			Name:    CompilerTemplateDepthKey,
			Aliases: []string{"templates", "template-depth", "templatedepth", "td"},
			Usage:   "max template depth for the compiler",
		},

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_COMPILER_CLONE_IMAGE", "COMPILER_CLONE_IMAGE"),
			Name:    CompilerCloneImageKey,
			Aliases: []string{"clone", "clone-image", "cloneimage", "ci"},
			Usage:   "image used to clone the repository for the compiler",
		},

		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_COMPILER_STARLARK_EXEC_LIMIT", "COMPILER_STARLARK_EXEC_LIMIT"),
			Name:    CompilerStarlarkExecLimitKey,
			Aliases: []string{"starlark-exec-limit", "starklark-limit", "starlarklimit", "sel"},
			Usage:   "max starlark execution limit for the compiler",
		},

		// SCM Flags

		&cli.StringMapFlag{
			Name:    SCMRepoRolesMapKey,
			Usage:   "map of SCM roles to Vela permissions for repositories",
			Sources: cli.EnvVars("VELA_SCM_REPO_ROLES_MAP", "SCM_REPO_ROLES_MAP"),
			Action: func(_ context.Context, _ *cli.Command, v map[string]string) error {
				return util.ValidateRoleMap(v, "repo")
			},
		},
		&cli.StringMapFlag{
			Name:    SCMOrgRolesMapKey,
			Usage:   "map of SCM roles to Vela permissions for organizations",
			Sources: cli.EnvVars("VELA_SCM_ORG_ROLES_MAP", "SCM_ORG_ROLES_MAP"),
			Action: func(_ context.Context, _ *cli.Command, v map[string]string) error {
				return util.ValidateRoleMap(v, "org")
			},
		},
		&cli.StringMapFlag{
			Name:    SCMTeamRolesMapKey,
			Usage:   "map of SCM roles to Vela permissions for teams",
			Sources: cli.EnvVars("VELA_SCM_TEAM_ROLES_MAP", "SCM_TEAM_ROLES_MAP"),
			Action: func(_ context.Context, _ *cli.Command, v map[string]string) error {
				return util.ValidateRoleMap(v, "team")
			},
		},

		// Misc Flags

		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_REPO_ALLOWLIST_ADD_REPOS", "REPO_ALLOWLIST_ADD_REPOS"),
			Name:    RepoAllowlistAddKey,
			Aliases: []string{"repo", "repos", "ral"},
			Usage:   "the list of repositories to add to the list of all those permitted to use Vela",
		},

		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_REPO_ALLOWLIST_DROP_REPOS", "REPO_ALLOWLIST_DROP_REPOS"),
			Name:    RepoAllowlistDropKey,
			Usage:   "the list of repositories to drop from the list of all those permitted to use Vela",
		},

		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_SCHEDULE_ALLOWLIST_ADD_REPOS", "SCHEDULE_ALLOWLIST_ADD_REPOS"),
			Name:    ScheduleAllowAddlistKey,
			Aliases: []string{"schedule", "schedules", "sal"},
			Usage:   "the list of repositories to add to the list of all those permitted to use schedules in Vela",
		},

		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_SCHEDULE_ALLOWLIST_DROP_REPOS", "SCHEDULE_ALLOWLIST_DROP_REPOS"),
			Name:    ScheduleAllowDroplistKey,
			Usage:   "the list of repositories to drop from the list of all those permitted to use schedules in Vela",
		},

		&cli.Int32Flag{
			Sources: cli.EnvVars("VELA_MAX_DASHBOARD_REPOS", "MAX_DASHBOARD_REPOS"),
			Name:    MaxDashboardReposKey,
			Usage:   "the maximum number of repositories for any dashboard",
		},

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_FILE", "SETTINGS_FILE"),
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "provide a file to update settings",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "SETTINGS_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update settings to change the compiler clone image to target/vela-git:latest.
    $ {{.FullName}} --compiler.clone-image target/vela-git:latest
  2. Update settings to change the compiler template depth to 2.
    $ {{.FullName}} --compiler.template-depth 2
  3. Update settings to change the compiler starlark exec limit to 5.
    $ {{.FullName}} --compiler.starlark-exec-limit 5
  4. Update settings with additional queue routes.
    $ {{.FullName}} --queue.add-route large --queue.add-route small
  5. Update settings by dropping queue routes.
    $ {{.FullName}} --queue.drop-route large --queue.drop-route small
  6. Update settings with additional repos permitted to use Vela (patterns containing * wildcards must be wrapped in quotes on the commandline).
    $ {{.FullName}} --add-repo octocat/hello-world --add-repo 'octocat/*'
  7. Update settings with additional repos permitted to use schedules in Vela (patterns containing * wildcards must be wrapped in quotes on the commandline).
    $ {{.FullName}} --add-schedule octocat/hello-world --schedule 'octocat/*'
  8. Update settings from a file.
    $ {{.FullName}} --file settings.yml
DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/settings/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to modify settings.
func update(_ context.Context, c *cli.Command) error {
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

	// create the settings configuration
	s := &settings.Config{
		Action: internal.ActionUpdate,
		Output: c.String(internal.FlagOutput),
		File:   c.String("file"),
		Queue: settings.Queue{
			AddRoutes:  c.StringSlice(QueueRouteAddKey),
			DropRoutes: c.StringSlice(QueueRouteDropKey),
		},
		Compiler:                   settings.Compiler{},
		SCM:                        settings.SCM{},
		RepoAllowlistAddRepos:      c.StringSlice(RepoAllowlistAddKey),
		RepoAllowlistDropRepos:     c.StringSlice(RepoAllowlistDropKey),
		ScheduleAllowlistAddRepos:  c.StringSlice(ScheduleAllowAddlistKey),
		ScheduleAllowlistDropRepos: c.StringSlice(ScheduleAllowDroplistKey),
		MaxDashboardRepos:          c.Int32(MaxDashboardReposKey),
	}

	// compiler
	if c.IsSet(CompilerCloneImageKey) {
		s.CloneImage = vela.String(c.String(CompilerCloneImageKey))
	}

	if c.IsSet(CompilerTemplateDepthKey) {
		s.TemplateDepth = vela.Int(c.Int(CompilerTemplateDepthKey))
	}

	if c.IsSet(CompilerStarlarkExecLimitKey) {
		s.StarlarkExecLimit = vela.Int64(c.Int64(CompilerStarlarkExecLimitKey))
	}

	// scm
	if c.IsSet(SCMRepoRolesMapKey) {
		s.RepoRoleMap = c.StringMap(SCMRepoRolesMapKey)
	}

	if c.IsSet(SCMOrgRolesMapKey) {
		s.OrgRoleMap = c.StringMap(SCMOrgRolesMapKey)
	}

	if c.IsSet(SCMTeamRolesMapKey) {
		s.TeamRoleMap = c.StringMap(SCMTeamRolesMapKey)
	}

	// validate settings configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/settings?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// check if file is provided
	if len(s.File) > 0 {
		return s.UpdateFromFile(client)
	}

	// execute the update call for the settings configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/settings?tab=doc#Config.Update
	return s.Update(client)
}
