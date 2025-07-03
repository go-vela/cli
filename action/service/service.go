// SPDX-License-Identifier: Apache-2.0

package service

import "github.com/go-vela/cli/internal/output"

// Config represents the configuration necessary
// to perform service related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Build   int64
	Number  int32
	Page    int
	PerPage int
	Output  string
	Color   output.ColorOptions
}
