// SPDX-License-Identifier: Apache-2.0

package output

// output drivers.
const (
	// DriverStdout defines the driver type
	// when outputting in stdout format.
	DriverStdout = "stdout"

	// DriverStderr defines the driver type
	// when outputting in stderr format.
	DriverStderr = "stderr"

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
