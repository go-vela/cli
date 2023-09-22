// SPDX-License-Identifier: Apache-2.0

package completion

// Config represents the configuration necessary
// to perform completion related requests with Vela.
type Config struct {
	Action string
	Bash   bool
	Zsh    bool
}
