// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"sort"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/types/library"

	"github.com/gosuri/uitable"

	"github.com/sirupsen/logrus"
)

// table is a helper function to output the
// provided pipelines in a table format with
// a specific set of fields displayed.
func table(pipelines *[]library.Pipeline) error {
	logrus.Debug("creating table for list of pipelines")

	// create a new table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#New
	table := uitable.New()

	// set column width for table to 50
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.Wrap = true

	logrus.Trace("adding headers to pipeline table")

	// set of pipeline fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("COMMIT", "REF", "TYPE", "VERSION", "STAGES", "STEPS")

	// iterate through all pipelines in the list
	for _, p := range reverse(*pipelines) {
		logrus.Tracef("adding pipeline %s to pipeline table", p.GetCommit())

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(p.GetCommit(), p.GetRef(), p.GetType(), p.GetVersion(), p.GetStages(), p.GetSteps())
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided pipelines in a wide table format with
// a specific set of fields displayed.
func wideTable(pipelines *[]library.Pipeline) error {
	logrus.Debug("creating wide table for list of pipelines")

	// create new wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#New
	table := uitable.New()

	// set column width for wide table to 200
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.Wrap = true

	logrus.Trace("adding headers to wide pipeline table")

	// set of pipeline fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("COMMIT", "REF", "TYPE", "VERSION", "EXTERNAL_SECRETS", "INTERNAL_SECRETS", "SERVICES", "STAGES", "STEPS", "TEMPLATES")

	// iterate through all pipelines in the list
	for _, p := range reverse(*pipelines) {
		logrus.Tracef("adding pipeline %s to pipeline table", p.GetCommit())

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(p.GetCommit(), p.GetRef(), p.GetType(), p.GetVersion(), p.GetExternalSecrets(), p.GetInternalSecrets(), p.GetServices(), p.GetStages(), p.GetSteps(), p.GetTemplates())
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// reverse is a helper function to sort the pipelines
// based off the pipeline number and then flip the
// order they get displayed in.
func reverse(p []library.Pipeline) []library.Pipeline {
	// sort the list of pipelines based off the pipeline id
	sort.SliceStable(p, func(i, j int) bool {
		return p[i].GetID() < p[j].GetID()
	})

	return p
}
