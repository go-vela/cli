// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

// Config represents the configuration necessary
// to perform worker related requests with Vela.
type Config struct {
	Action            string
	Address           string
	Hostname          string
	Active            bool
	Routes            []string
	BuildLimit        int64
	RegistrationToken bool
	Page              int
	PerPage           int
	Output            string
}
