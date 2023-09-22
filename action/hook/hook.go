// SPDX-License-Identifier: Apache-2.0

package hook

// Config represents the configuration necessary
// to perform hook related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Number  int
	Page    int
	PerPage int
	Output  string
}
