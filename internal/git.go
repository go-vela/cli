// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package internal

import (
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	giturls "github.com/whilp/git-urls"
)

// SetGitConfigContext attempts to set the org and repo
// based on the .git/ directory, provided the user has
// the config flag of gitsync set to true.
func SetGitConfigContext(c *cli.Context) {
	// check to see if config and command allow for
	// automatic setting of org and repo.
	if c.String(FlagGitSync) == "true" &&
		c.String(FlagOrg) == "" &&
		c.String(FlagRepo) == "" {
		// set the org
		logrus.Trace("attempting to set org from .git/")
		err := c.Set(FlagOrg, GetGitConfigOrg("./"))

		if err != nil {
			logrus.Debug("failed to set org in context")
		}

		// set the repo
		logrus.Trace("attempting to set repo from .git/")
		err = c.Set(FlagRepo, GetGitConfigRepo("./"))

		if err != nil {
			logrus.Debug("failed to set repo in context")
		}
	}
}

// GetGitConfigOrg opens the git repository, fetches
// the remote origin url, and parses the url to find
// the org of the current working directory.
func GetGitConfigOrg(path string) string {
	// open repository
	r, err := git.PlainOpen(path)

	// on failure, return empty string to allow for
	// potential manual setting of org to process
	if err != nil {
		return ""
	}

	// fetch remote origin
	origin, err := r.Remote("origin")
	if err != nil {
		return ""
	}
	url, err := giturls.Parse(origin.Config().URLs[0])
	if err != nil {
		return ""
	}

	// path is :org/:repo.git - get org
	org := strings.Split(url.Path, "/")[0]

	return org
}

// GetGitConfigRepo opens the git repository, fetches
// the remote origin url, and parses the url to find
// the repo of the current working directory.
func GetGitConfigRepo(path string) string {
	// open repository
	r, err := git.PlainOpen(path)

	// on failure, return empty string to allow for
	// potential manual setting of org to process
	if err != nil {
		return ""
	}

	// fetch remote origin
	origin, err := r.Remote("origin")
	if err != nil {
		return ""
	}
	url, err := giturls.Parse(origin.Config().URLs[0])
	if err != nil {
		return ""
	}

	// path is :org/:repo.git - get repo
	repoName := strings.Split(strings.Split(url.Path, "/")[1], ".git")[0]
	return repoName
}
