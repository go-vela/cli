// SPDX-License-Identifier: Apache-2.0

package config

// Config represents the configuration necessary
// to perform config related quests with Vela.
type Config struct {
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
	GitHub       *GitHub
	UpdateFlags  map[string]string
	RemoveFlags  []string
	Output       string
}
