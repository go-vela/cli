// SPDX-License-Identifier: Apache-2.0

package dashboard

// Config represents the configuration necessary
// to perform dashboard related requests with Vela.
type Config struct {
	Action      string
	Name        string
	ID          string
	AddRepos    []string
	TargetRepos []string
	DropRepos   []string
	Branches    []string
	Events      []string
	AddAdmins   []string
	DropAdmins  []string
	Full        bool
	Output      string
}
