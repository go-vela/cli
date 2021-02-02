// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

// output drivers.
const (
	// DriverStdout defines the driver type
	// when outputting in stdout format.
	DriverStdout = "stdout"

	// DriverDump defines the driver type
	// when outputting in dump format.
	DriverDump = "dump"

	// DriverJSON defines the driver type
	// when outputting in JSON format.
	DriverJSON = "json"

	// DriverRawJSON defines the driver type
	// when outputting in raw JSON format.
	DriverRawJSON = "rawjson"

	// DriverSpew defines the driver type
	// when outputting in github.com/davecgh/go-spew/spew format.
	DriverSpew = "spew"

	// DriverYAML defines the driver type
	// when outputting in YAML format.
	DriverYAML = "yaml"
)
