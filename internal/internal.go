// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package internal

// For more information:
//
//   * https://golang.org/doc/go1.4#internalpackages
//   * https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit

// API flag keys.
const (
	// FlagAPIAddress defines the key for the
	// flag when setting the API address.
	FlagAPIAddress = "api.addr"

	// FlagAPIToken defines the key for the
	// flag when setting the API token.
	FlagAPIToken = "api.token"

	// FlagAPIAccessToken defines the key for
	// the flag when setting the API access token.
	FlagAPIAccessToken = "api.token.access"

	// FlagAPIRefreshToken defines the key for
	// for the flag when setting the API
	// refresh token.
	// nolint:gosec // false negative - not a real token
	FlagAPIRefreshToken = "api.token.refresh"

	// FlagAPIVersion defines the key for the
	// flag when setting the API version.
	FlagAPIVersion = "api.version"
)

// build flag keys.
const (
	// FlagBuild defines the key for the
	// flag when setting the build.
	FlagBuild = "build"
)

// compiler flag keys.
const (
	// FlagCompilerGitHubToken defines the key for the
	// flag when setting the compiler github token.
	// nolint:gosec // ignoring since this is a constant for a user passed token
	FlagCompilerGitHubToken = "compiler.github.token"

	// FlagCompilerGitHubURL defines the key for the
	// flag when setting the compiler github url.
	FlagCompilerGitHubURL = "compiler.github.url"

	// FlagCompilerGithubDriver defines the key for the
	// flag when setting the compiler github driver.
	FlagCompilerGithubDriver = "compiler.github.driver"
)

// generic flag keys.
const (
	// FlagConfig defines the key for the
	// flag when setting the config.
	FlagConfig = "config"

	// FlagOutput defines the key for the
	// flag when setting the output.
	FlagOutput = "output"
)

// log flag keys.
const (
	// FlagLogLevel defines the key for the
	// flag when setting the log level.
	FlagLogLevel = "log.level"
)

// no git flag keys.
const (
	// FlagNoGit defines the key for the
	// flag when setting the no-git status.
	FlagNoGit = "no-git"
)

// pagination flag keys.
const (
	// FlagPage defines the key for the
	// flag when setting the page.
	FlagPage = "page"

	// FlagPerPage defines the key for the
	// flag when setting the results per page.
	FlagPerPage = "per.page"
)

// repository flag keys.
const (
	// FlagOrg defines the key for the
	// flag when setting the org.
	FlagOrg = "org"

	// FlagRepo defines the key for the
	// flag when setting the repo.
	FlagRepo = "repo"
)

// secret flag keys.
const (
	// FlagSecretEngine defines the key for the
	// flag when setting the secret engine.
	FlagSecretEngine = "secret.engine"

	// FlagSecretType defines the key for the
	// flag when setting the secret type.
	FlagSecretType = "secret.type"
)

// service flag keys.
const (
	// FlagService defines the key for the
	// flag when setting the service.
	FlagService = "service"
)

// step flag keys.
const (
	// FlagStep defines the key for the
	// flag when setting the step.
	FlagStep = "step"
)

// list of defined CLI actions.
const (
	// ActionAdd defines the action for creating a resource.
	ActionAdd = "add"

	// ActionCancel defines the action for canceling of a resource.
	ActionCancel = "cancel"

	// ActionChown defines the action for changing ownership of a resource.
	ActionChown = "chown"

	// ActionCompile defines the action for compiling a resource.
	ActionCompile = "compile"

	// ActionExec defines the action for executing a resource.
	ActionExec = "exec"

	// ActionExpand defines the action for expanding a resource.
	ActionExpand = "expand"

	// ActionGenerate defines the action for producing a resource.
	ActionGenerate = "generate"

	// ActionGet defines the action for getting a list of resources.
	ActionGet = "get"

	// ActionLoad defines the action for loading a resource.
	ActionLoad = "load"

	// ActionRemove defines the action for deleting a resource.
	ActionRemove = "remove"

	// ActionRepair defines the action for repairing a resource.
	ActionRepair = "repair"

	// ActionRestart defines the action for restarting a resource.
	ActionRestart = "restart"

	// ActionSync defines the action for syncing a resource with SCM.
	ActionSync = "sync"

	// ActionSyncAll defines the action for syncing all org resources with SCM.
	ActionSyncAll = "syncAll"

	// ActionUpdate defines the action for modifying a resource.
	ActionUpdate = "update"

	// ActionValidate defines the action for validating a resource.
	ActionValidate = "validate"

	// ActionView defines the action for inspecting a resource.
	ActionView = "view"
)
