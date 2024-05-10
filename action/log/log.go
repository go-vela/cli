// SPDX-License-Identifier: Apache-2.0

package log

import "github.com/go-vela/cli/internal/output"

// Config represents the configuration necessary
// to perform log related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Build   int
	Page    int
	PerPage int
	Service int
	Step    int
	Output  string
	Color   output.ColorOptions
}
