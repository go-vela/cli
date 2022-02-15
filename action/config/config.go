// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

// Config represents the configuration necessary
// to perform config related quests with Vela.
type Config struct {
	UpdateFlags  map[string]string
	GitHub       *GitHub
	Action       string
	File         string
	Addr         string
	Token        string
	AccessToken  string
	RefreshToken string
	Version      string
	LogLevel     string
	NoGit        string
	Org          string
	Repo         string
	Engine       string
	Type         string
	Output       string
	RemoveFlags  []string
}
