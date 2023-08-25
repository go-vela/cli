// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package schedule

// Config represents the configuration necessary
// to perform schedule related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Active  bool
	Name    string
	Entry   string
	Page    int
	PerPage int
	Output  string
	Branch  string
}
