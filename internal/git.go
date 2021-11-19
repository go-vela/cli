// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package internal

import (
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	giturls "github.com/whilp/git-urls"
)

// SetGitConfigContext attempts to set the org and repo
// based on the .git/ directory, provided the user has
// the config flag of no-git set to true.
func SetGitConfigContext(c *cli.Context) {
	// check to see if config and command allow for
	// automatic setting of org and repo.
	nogit, err := strconv.ParseBool(c.String(FlagNoGit))
	if err != nil {
		logrus.Debug("Invalid bool value for flag no-git")
		return
	}
	if !nogit &&
		c.String(FlagOrg) == "" &&
		c.String(FlagRepo) == "" {
		// set the org
		logrus.Trace("attempting to set org from .git/")
		err = c.Set(FlagOrg, GetGitConfigOrg("./"))

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
	// nolint: gomnd // ignore magic number
	splitOrg := strings.SplitN(url.Path, "/", 2)

	// check if url path is expected format
	// nolint: gomnd // ignore magic number
	if len(splitOrg) != 2 {
		logrus.Debug("Invalid remote origin url -- please specify org and repo")
		return ""
	}

	// path is :org/:repo.git - get org
	org := splitOrg[0]

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

	// nolint: gomnd // ignore magic number
	splitRepo := strings.SplitN(url.Path, "/", 2)

	// check if url path is expected format
	// nolint: gomnd // ignore magic number
	if len(splitRepo) != 2 {
		logrus.Debug("Invalid remote origin url -- please specify org and repo")
		return ""
	}

	// nolint: gomnd // ignore magic number
	splitDotGit := strings.SplitN(splitRepo[1], ".git", 2)

	// check if repo name is expected format
	// nolint: gomnd // ignore magic number
	if len(splitDotGit) != 2 {
		logrus.Debug("Invalid remote origin url -- please specify org and repo")
		return ""
	}

	repoName := splitDotGit[0]
	return repoName
}
