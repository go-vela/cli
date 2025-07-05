// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"go.yaml.in/yaml/v3"
)

// create filesystem based on the operating system
//
// https://godoc.org/github.com/spf13/afero#NewOsFs
var appFS = afero.NewOsFs()

// Generate produces a pipeline based off the provided configuration.
func (c *Config) Generate() error {
	logrus.Debug("executing generate for pipeline configuration")

	// create the pipeline file content
	pipeline := steps(c.Type)

	// check if stages were enabled for the provided configuration
	if c.Stages {
		pipeline = stages(c.Type)
	}

	logrus.Trace("creating file content from pipeline")

	// create output for pipeline file
	out, err := yaml.Marshal(pipeline)
	if err != nil {
		return err
	}

	// use custom filesystem which enables us to test
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero
	a := &afero.Afero{
		Fs: appFS,
	}

	// send Filesystem call to capture base directory path
	base, err := os.Getwd()
	if err != nil {
		return err
	}

	// create full path for pipeline file
	path := filepath.Join(base, c.File)

	// check if custom path was provided for pipeline file
	if len(c.Path) > 0 {
		// create custom full path for pipeline file
		path = filepath.Join(c.Path, c.File)
	}

	logrus.Tracef("creating directory structure to %s", path)

	// send Filesystem call to create directory path for pipeline file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#OsFs.MkdirAll
	err = a.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return err
	}

	logrus.Tracef("writing file content to %s", path)

	// send Filesystem call to create pipeline file
	//
	// https://pkg.go.dev/github.com/spf13/afero?tab=doc#Afero.WriteFile
	return a.WriteFile(path, out, 0644)
}
