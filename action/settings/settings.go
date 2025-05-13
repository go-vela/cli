// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"github.com/go-vela/cli/internal/output"
)

// Config represents the configuration necessary
// to perform settings related requests with Vela.
type Config struct {
	Action string
	Output string
	File   string
	Compiler
	Queue
	SCM

	RepoAllowlist          *[]string
	RepoAllowlistAddRepos  []string
	RepoAllowlistDropRepos []string

	ScheduleAllowlist          *[]string
	ScheduleAllowlistAddRepos  []string
	ScheduleAllowlistDropRepos []string

	MaxDashboardRepos int32

	Color output.ColorOptions
}

// Compiler represents the compiler configurations used
// to modify the compiler settings for Vela.
type Compiler struct {
	CloneImage        *string
	TemplateDepth     *int
	StarlarkExecLimit *int64
}

// Queue represents the compiler configurations used
// to modify the queue settings for Vela.
type Queue struct {
	Routes     *[]string
	AddRoutes  []string
	DropRoutes []string
}

// SCM represents the SCM configurations used
// to modify the SCM settings for Vela.
type SCM struct {
	RepoRoleMap map[string]string
	OrgRoleMap  map[string]string
	TeamRoleMap map[string]string
}
