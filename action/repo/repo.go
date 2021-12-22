// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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
	BuildLimit   int
	Timeout      int64
	Counter      int
	Private      bool
	Trusted      bool
	Active       bool
	Events       []string
	PipelineType string
	Page         int
	PerPage      int
	Output       string
}
