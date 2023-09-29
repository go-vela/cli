// SPDX-License-Identifier: Apache-2.0

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

	if !nogit && c.String(FlagOrg) == "" && c.String(FlagRepo) == "" {
		logrus.Trace("attempting to set org from .git/")

		// set the org
		err = c.Set(FlagOrg, GetGitConfigOrg("./"))
		if err != nil {
			logrus.Debug("failed to set org in context")
		}

		logrus.Trace("attempting to set repo from .git/")

		// set the repo
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

	urlPath := strings.TrimPrefix(url.Path, "/")

	splitOrg := strings.SplitN(urlPath, "/", 2)
	if len(splitOrg) != 2 {
		logrus.Debug("Invalid remote origin url -- please specify org and repo")
		return ""
	}

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

	urlPath := strings.TrimPrefix(url.Path, "/")

	splitRepo := strings.SplitN(urlPath, "/", 2)
	if len(splitRepo) != 2 {
		logrus.Debug("Invalid remote origin url -- please specify org and repo")
		return ""
	}

	repo := splitRepo[1]

	// chop off .git at the end of the repo name if it exists
	splitDotGit := strings.SplitN(repo, ".git", 2)

	repoName := splitDotGit[0]

	return repoName
}
