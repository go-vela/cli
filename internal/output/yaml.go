// SPDX-License-Identifier: Apache-2.0

package output

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.yaml.in/yaml/v3"
)

// YAML parses the provided input and
// renders the parsed input in YAML
// before outputting it to stdout.
func YAML(_input interface{}, colorOpts ColorOptions) error {
	logrus.Debugf("creating output with %s driver", DriverYAML)

	// validate the input provided
	err := validate(DriverYAML, _input)
	if err != nil {
		return err
	}

	// marshal the input into YAML
	output, err := yaml.Marshal(_input)
	if err != nil {
		return err
	}

	logrus.Tracef("sending output to stdout with %s driver", DriverYAML)

	// attempt to highlight the output
	// returns the input and logs a warning on failure
	strOutput := Highlight(string(output), "yaml", colorOpts)

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, strOutput)

	return nil
}
