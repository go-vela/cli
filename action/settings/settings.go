// SPDX-License-Identifier: Apache-2.0

package settings

// Config represents the configuration necessary
// to perform settings related requests with Vela.
type Config struct {
	Action string
	Compiler
	Queue
	RepoAllowlist     *[]string
	ScheduleAllowlist *[]string
	Output            string
}

type Compiler struct {
	CloneImage        *string
	TemplateDepth     *int
	StarlarkExecLimit *uint64
}

type Queue struct {
	Routes *[]string
}
