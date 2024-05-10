// SPDX-License-Identifier: Apache-2.0

package output

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
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

	strOutput := string(output)

	if colorOpts.Enabled {
		buf := new(bytes.Buffer)
		err = quick.Highlight(buf, strOutput, "yaml", colorOpts.Format, colorOpts.Theme)
		if err == nil {
			strOutput = buf.String()
		} else {
			logrus.Warnf("unable to highlight yaml output: %v", err)
		}
	}

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, strOutput)

	return nil
}
