package cmd

import (
	"github.com/go-vela/cli/cmd/completion"
	"github.com/urfave/cli/v2"
)

var completionCmd = cli.Command{
	Name:        "completion",
	Category:    "User experience",
	Aliases:     []string{"c"},
	Description: "Use this command to enable vela auto completion in your shell.",
	Usage:       "Enable vela auto completion for the current session. Supports bash & zsh.",
	Subcommands: []*cli.Command{
		&completion.BashCmd,
		&completion.ZSHCmd,
	},
}
