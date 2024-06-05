// SPDX-License-Identifier: Apache-2.0

package settings

import "github.com/go-vela/cli/internal/output"

// Config represents the configuration necessary
// to perform settings related requests with Vela.
type Config struct {
	Action string
	Output string
	File   string
	Compiler
	Queue

	RepoAllowlist          *[]string
	RepoAllowlistAddRepos  []string
	RepoAllowlistDropRepos []string

	ScheduleAllowlist          *[]string
	ScheduleAllowlistAddRepos  []string
	ScheduleAllowlistDropRepos []string

	Color output.ColorOptions
}

// Compiler represents the compiler configurations used
// to modify the compiler settings for Vela.
type Compiler struct {
	CloneImage        *string
	TemplateDepth     *int
	StarlarkExecLimit *uint64
}

// Queue represents the compiler configurations used
// to modify the queue settings for Vela.
type Queue struct {
	Routes     *[]string
	AddRoutes  []string
	DropRoutes []string
}
