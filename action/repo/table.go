// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"strings"

	"github.com/gosuri/uitable"
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/types/library"
)

// table is a helper function to output the
// provided repos in a table format with
// a specific set of fields displayed.
func table(repos *[]library.Repo) error {
	logrus.Debug("creating table for list of repos")

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

	logrus.Trace("adding headers to repo table")

	// set of repository fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("ORG/REPO", "ACTIVE", "EVENTS", "VISIBILITY", "BRANCH")

	// iterate through all repos in the list
	for _, r := range *repos {
		logrus.Tracef("adding repo %s to repo table", r.GetFullName())

		
		e := strings.Join(r.AllowEvents.List(), ",")

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(r.GetFullName(), r.GetActive(), e, r.GetVisibility(), r.GetBranch())
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided repos in a wide table format with
// a specific set of fields displayed.
func wideTable(repos *[]library.Repo) error {
	logrus.Debug("creating wide table for list of repos")

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

	logrus.Trace("adding headers to wide repo table")

	// set of repository fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("ORG/REPO", "ACTIVE", "EVENTS", "VISIBILITY", "BRANCH", "REMOTE")

	// iterate through all repos in the list
	for _, r := range *repos {
		logrus.Tracef("adding repo %s to wide repo table", r.GetFullName())

		
		e := strings.Join(r.AllowEvents.List(), ",")

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(r.GetFullName(), r.GetActive(), e, r.GetVisibility(), r.GetBranch(), r.GetLink())
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}
