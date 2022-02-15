// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

// Config represents the configuration necessary
// to perform step related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Output  string
	Build   int
	Number  int
	Page    int
	PerPage int
}
