// SPDX-License-Identifier: Apache-2.0

package user

import (
	"github.com/go-vela/cli/internal/output"
	api "github.com/go-vela/server/api/types"
)

// Config represents the configuration necessary
// to perform user related requests with Vela.
type Config struct {
	Name           string
	AddFavorites   []string
	DropFavorites  []string
	AddDashboards  []string
	DropDashboards []string
	Output         string
	Color          output.ColorOptions
}

func outputUser(user *api.User, c *Config) error {
	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the user in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(user)
	case output.DriverJSON:
		// output the user in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(user, c.Color)
	case output.DriverSpew:
		// output the user in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(user)
	case output.DriverYAML:
		// output the user in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(user, c.Color)
	default:
		// output the user in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(user)
	}
}
