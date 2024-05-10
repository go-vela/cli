// SPDX-License-Identifier: Apache-2.0

package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// JSON parses the provided input and
// renders the parsed input in pretty
// JSON before outputting it to stdout.
func JSON(_input interface{}, colorOpts ColorOptions) error {
	logrus.Debugf("creating output with %s driver", DriverJSON)

	// validate the input provided
	err := validate(DriverJSON, _input)
	if err != nil {
		return err
	}

	// marshal the input into pretty JSON
	output, err := json.MarshalIndent(_input, "", "    ")
	if err != nil {
		return err
	}

	logrus.Tracef("sending output to stdout with %s driver", DriverJSON)

	// attempt to highlight the output
	// returns the input and logs a warning on failure
	strOutput := Highlight(string(output), "json", colorOpts)

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, strOutput)

	return nil
}
