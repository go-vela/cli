// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package internal

// For more information:
//
//   * https://golang.org/doc/go1.4#internalpackages
//   * https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit

// API flag keys
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

// build flag keys
const (
	// FlagBuild defines the key for the
	// flag when setting the build.
	FlagBuild = "build"
)

// generic flag keys
const (
	// FlagConfig defines the key for the
	// flag when setting the config.
	FlagConfig = "config"

	// FlagOutput defines the key for the
	// flag when setting the output.
	FlagOutput = "output"
)

// log flag keys
const (
	// FlagLogLevel defines the key for the
	// flag when setting the log level.
	FlagLogLevel = "log.level"
)

// pagination flag keys
const (
	// FlagPage defines the key for the
	// flag when setting the page.
	FlagPage = "page"

	// FlagPerPage defines the key for the
	// flag when setting the results per page.
	FlagPerPage = "per.page"
)

// repository flag keys
const (
	// FlagOrg defines the key for the
	// flag when setting the org.
	FlagOrg = "org"

	// FlagRepo defines the key for the
	// flag when setting the repo.
	FlagRepo = "repo"
)

// secret flag keys
const (
	// FlagSecretEngine defines the key for the
	// flag when setting the secret engine.
	FlagSecretEngine = "secret.engine"

	// FlagSecretType defines the key for the
	// flag when setting the secret type.
	FlagSecretType = "secret.type"
)

// service flag keys
const (
	// FlagService defines the key for the
	// flag when setting the service.
	FlagService = "service"
)

// step flag keys
const (
	// FlagStep defines the key for the
	// flag when setting the step.
	FlagStep = "step"
)
