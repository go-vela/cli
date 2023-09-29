// SPDX-License-Identifier: Apache-2.0

package build

// Config represents the configuration necessary
// to perform build related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Number  int
	Event   string
	Status  string
	Branch  string
	Before  int64
	After   int64
	Page    int
	PerPage int
	Output  string
}
