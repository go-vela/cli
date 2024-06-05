// SPDX-License-Identifier: Apache-2.0

package schedule

import "github.com/go-vela/cli/internal/output"

// Config represents the configuration necessary
// to perform schedule related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Active  bool
	Name    string
	Entry   string
	Page    int
	PerPage int
	Output  string
	Branch  string
	Color   output.ColorOptions
}
