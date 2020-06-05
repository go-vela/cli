package hook

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/go-vela/cli/util"
	"github.com/go-vela/types/library"
	"github.com/gosuri/uitable"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"time"
)

func validate(c *cli.Context) error {

	if len(c.String("org")) == 0 {
		return util.InvalidCommand("org")
	}

	if len(c.String("repo")) == 0 {
		return util.InvalidCommand("repo")
	}

	return nil
}

// Helper function to print hook in the specified format (if none specified, prints in table format)
func PrintOutput(format string, hooks ...library.Hook) error {
	switch format {

	case "yaml":
		output, err := yaml.Marshal(hooks)
		if err != nil {
			return err
		}
		fmt.Println(string(output))

	case "json":
		output, err := json.MarshalIndent(hooks, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(output))

	case "wide":
		table := uitable.New()
		table.MaxColWidth = 100
		table.Wrap = true

		// Add headers
		table.AddRow("HOOK_NUMBER", "HOOK_STATUS", "EVENT", "BRANCH", "CREATED", "ERROR")
		for _, hook := range hooks {
			// Add data
			table.AddRow(*hook.Number, *hook.Status, *hook.Event, *hook.Branch, humanize.Time(time.Unix(*hook.Created, 0)), *hook.Error)
		}
		fmt.Println(table)

	default:
		table := uitable.New()
		table.MaxColWidth = 100
		table.Wrap = true

		// Add headers
		table.AddRow("HOOK_NUMBER", "HOOK_STATUS", "EVENT", "BRANCH")
		for _, hook := range hooks {
			// Add data
			table.AddRow(*hook.Number, *hook.Status, *hook.Event, *hook.Branch)
		}
		fmt.Println(table)
	}

	return nil
}
