// SPDX-License-Identifier: Apache-2.0

package step

// Config represents the configuration necessary
// to perform step related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Build   int
	Number  int
	Page    int
	PerPage int
	Output  string
}
