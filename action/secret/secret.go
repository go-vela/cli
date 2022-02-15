// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

// Config represents the configuration necessary
// to perform secret related requests with Vela.
type Config struct {
	Action       string
	Engine       string
	Type         string
	Org          string
	Repo         string
	Team         string
	Name         string
	Value        string
	File         string
	Output       string
	Images       []string
	Events       []string
	Page         int
	PerPage      int
	AllowCommand bool
}

// setValue is a helper function to check if the value
// was provided directly as a flag or via input from file.
func (c *Config) setValue() error {
	logrus.Debugf("capturing value for secret %s", c.Name)

	// check if the '@' character was provided signaling
	// we should capture the value from a file
	if strings.HasPrefix(c.Value, "@") {
		// capture the original path to the file by trimming the '@' character
		path := strings.TrimPrefix(c.Value, "@")

		logrus.Tracef("reading contents from %s", path)

		// capture the contents from the file to be added as a secret value
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("unable to read file %s: %w", path, err)
		}

		// set the secret value to the contents from the file
		c.Value = string(data)
	}

	return nil
}
