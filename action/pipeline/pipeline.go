// SPDX-License-Identifier: Apache-2.0

package pipeline

import "github.com/go-vela/cli/internal/output"

// Config represents the configuration necessary
// to perform pipeline related requests with Vela.
type Config struct {
	Action           string
	Branch           string
	Comment          string
	Event            string
	Status           string
	Tag              string
	Target           string
	Org              string
	Repo             string
	Ref              string
	File             string
	FileChangeset    []string
	Path             string
	Type             string
	Stages           bool
	TemplateFiles    []string
	Local            bool
	Remote           bool
	Volumes          []string
	PrivilegedImages []string
	Page             int
	PerPage          int
	Output           string
	Color            output.ColorOptions
	PipelineType     string
}
