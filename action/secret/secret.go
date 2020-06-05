// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

// Config represents the configuration necessary
// to perform secret related quests with Vela.
type Config struct {
	Action       string
	Engine       string
	Type         string
	Org          string
	Repo         string
	Team         string
	Name         string
	Value        string
	Images       []string
	Events       []string
	AllowCommand bool
	File         string
	Page         int
	PerPage      int
	Output       string
}
