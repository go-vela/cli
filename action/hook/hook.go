// SPDX-License-Identifier: Apache-2.0

package hook

import "github.com/go-vela/cli/internal/output"

// Config represents the configuration necessary
// to perform hook related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Number  int64
	Page    int
	PerPage int
	Output  string
	Color   output.ColorOptions
}
