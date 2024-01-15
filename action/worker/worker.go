// SPDX-License-Identifier: Apache-2.0

package worker

// Config represents the configuration necessary
// to perform worker related requests with Vela.
type Config struct {
	Action            string
	Address           string
	CheckedInBefore   int64
	CheckedInAfter    int64
	Hostname          string
	Active            *bool
	Routes            []string
	BuildLimit        int64
	RegistrationToken bool
	Output            string
}
