// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/settings"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/sdk-go/vela"
)

const (
	QueueRouteKey                = "queue.routes"
	CompilerCloneImageKey        = "compiler.clone-image"
	CompilerTemplateDepthKey     = "compiler.template-depth"
	CompilerStarlarkExecLimitKey = "compiler.starlark-exec-limit"
	RepoAllowlistAddKey          = "add-repo"
	RepoAllowlistDropKey         = "drop-repo"
	ScheduleAllowAddlistKey      = "add-schedule"
	ScheduleAllowDroplistKey     = "drop-schedule"
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
			EnvVars: []string{"VELA_QUEUE_ROUTES", "SETTINGS_ROUTES", "QUEUE_ROUTES"},
			Name:    QueueRouteKey,
			Aliases: []string{"queue-route", "routes", "route", "r"},
			Usage:   "route assignment for the queue",
		},

		// Compiler Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_COMPILER_TEMPLATE_DEPTH", "VELA_TEMPLATE_DEPTH", "TEMPLATE_DEPTH"},
			Name:    CompilerTemplateDepthKey,
			Aliases: []string{"templates", "template-depth", "templatedepth", "td"},
			Usage:   "max template depth for the compiler",
		},

		&cli.StringFlag{
			EnvVars: []string{"VELA_COMPILER_CLONE_IMAGE", "COMPILER_CLONE_IMAGE"},
			Name:    CompilerCloneImageKey,
			Aliases: []string{"clone", "clone-image", "cloneimage", "ci"},
			Usage:   "image used to clone the repository for the compiler",
		},

		&cli.IntFlag{
			EnvVars: []string{"VELA_COMPILER_STARLARK_EXEC_LIMIT", "COMPILER_STARLARK_EXEC_LIMIT"},
			Name:    CompilerStarlarkExecLimitKey,
			Aliases: []string{"starlark-exec-limit", "starklark-limit", "starlarklimit", "sel"},
			Usage:   "max starlark execution limit for the compiler",
		},

		// Misc Flags

		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_REPO_ALLOWLIST_ADD_REPOS", "REPO_ALLOWLIST_ADD_REPOS"},
			Name:    RepoAllowlistAddKey,
			Aliases: []string{"repo", "repos", "ral"},
			Usage:   "the list of repositories to add to the list of all those permitted to use Vela",
		},

		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_REPO_ALLOWLIST_DROP_REPOS", "REPO_ALLOWLIST_DROP_REPOS"},
			Name:    RepoAllowlistDropKey,
			Usage:   "the list of repositories to drop from the list of all those permitted to use Vela",
		},

		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_SCHEDULE_ALLOWLIST_ADD_REPOS", "SCHEDULE_ALLOWLIST_ADD_REPOS"},
			Name:    ScheduleAllowAddlistKey,
			Aliases: []string{"schedule", "schedules", "sal"},
			Usage:   "the list of repositories to add to the list of all those permitted to use schedules in Vela",
		},

		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_SCHEDULE_ALLOWLIST_DROP_REPOS", "SCHEDULE_ALLOWLIST_DROP_REPOS"},
			Name:    ScheduleAllowDroplistKey,
			Usage:   "the list of repositories to drop from the list of all those permitted to use schedules in Vela",
		},

		&cli.StringFlag{
			EnvVars: []string{"VELA_FILE", "SECRET_FILE"},
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "provide a file to update settings",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SETTINGS_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update settings to change the compiler clone image to target/vela-git:latest.
    $ {{.HelpName}} --compiler.clone-image target/vela-git:latest
  2. Update settings to change the compiler template depth to 2.
    $ {{.HelpName}} --compiler.template-depth 2
  3. Update settings to change the compiler starlark exec limit to 5.
    $ {{.HelpName}} --compiler.starlark-exec-limit 5
  4. Update settings with additional queue routes.
    $ {{.HelpName}} --queue.route large --queue.route small
  5. Update settings with additional repos permitted to use Vela.
  	$ {{.HelpName}} --repo-allowlist octocat/hello-world --repo-allowlist octocat/*
  6. Update settings with additional repos permitted to use schedules in Vela.
	$ {{.HelpName}} --schedule-allowlist octocat/hello-world --schedule-allowlist octocat/*
DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/settings/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to modify settings.
func update(c *cli.Context) error {
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
		Action:                     internal.ActionUpdate,
		Output:                     c.String(internal.FlagOutput),
		File:                       c.String("file"),
		Queue:                      settings.Queue{},
		Compiler:                   settings.Compiler{},
		RepoAllowlistAddRepos:      c.StringSlice(RepoAllowlistAddKey),
		RepoAllowlistDropRepos:     c.StringSlice(RepoAllowlistDropKey),
		ScheduleAllowlistAddRepos:  c.StringSlice(ScheduleAllowAddlistKey),
		ScheduleAllowlistDropRepos: c.StringSlice(ScheduleAllowDroplistKey),
	}

	// queue
	if c.IsSet(QueueRouteKey) {
		s.Queue.Routes = vela.Strings(c.StringSlice(QueueRouteKey))
	}

	// compiler
	if c.IsSet(CompilerCloneImageKey) {
		s.Compiler.CloneImage = vela.String(c.String(CompilerCloneImageKey))
	}

	if c.IsSet(CompilerTemplateDepthKey) {
		s.Compiler.TemplateDepth = vela.Int(c.Int(CompilerTemplateDepthKey))
	}

	if c.IsSet(CompilerStarlarkExecLimitKey) {
		s.Compiler.StarlarkExecLimit = vela.UInt64(c.Uint64(CompilerStarlarkExecLimitKey))
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
