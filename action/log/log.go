// SPDX-License-Identifier: Apache-2.0

package log

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
}
