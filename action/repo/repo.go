// SPDX-License-Identifier: Apache-2.0

package repo

// Config represents the configuration necessary
// to perform repository related requests with Vela.
type Config struct {
	Action       string
	Org          string
	Name         string
	Branch       string
	Link         string
	Clone        string
	Visibility   string
	BuildLimit   int64
	Timeout      int64
	Counter      int
	Private      bool
	Trusted      bool
	Active       bool
	Events       []string
	PipelineType string
	ApproveBuild string
	Page         int
	PerPage      int
	Output       string
}

// Alternate constants for webhook events.
const (
	AlternateDeploy = "deploy"
	AlternatePull   = "pull"
)
