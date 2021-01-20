// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"sort"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/types/library"

	"github.com/gosuri/uitable"

	"github.com/sirupsen/logrus"
)

// table is a helper function to output the
// provided deployments in a table format with
// a specific set of fields displayed.
func table(deployments *[]library.Deployment) error {
	logrus.Debug("creating table for list of deployments")

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

	logrus.Trace("adding headers to deployment table")

	// set of deployment fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("ID", "TASK", "USER", "REF", "TARGET")

	// iterate through all deployments in the list
	for _, d := range reverse(*deployments) {
		logrus.Tracef("adding deployment %d to deployment table", d.GetID())

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(d.GetID(), d.GetTask(), d.GetUser(), d.GetRef(), d.GetTarget())
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided deployments in a wide table format with
// a specific set of fields displayed.
func wideTable(deployments *[]library.Deployment) error {
	logrus.Debug("creating wide table for list of deployments")

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

	logrus.Trace("adding headers to wide deployment table")

	// set of deployment fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("ID", "TASK", "USER", "REF", "TARGET", "COMMIT", "DESCRIPTION")

	// iterate through all deployments in the list
	for _, d := range reverse(*deployments) {
		logrus.Tracef("adding deployment %d to wide deployment table", d.GetID())

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(d.GetID(), d.GetTask(), d.GetUser(), d.GetRef(), d.GetTarget(), d.GetCommit(), d.GetDescription())
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// reverse is a helper function to sort the deployments
// based off the deployment number and then flip the
// order they get displayed in.
func reverse(d []library.Deployment) []library.Deployment {
	// sort the list of deployments based off the deployment id
	sort.SliceStable(d, func(i, j int) bool {
		return d[i].GetID() < d[j].GetID()
	})

	return d
}
