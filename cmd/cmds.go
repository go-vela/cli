// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/urfave/cli"
)

// Vela holds all of the top level commands in the CLI
var Vela = []cli.Command{

	addCmds,
	updateCmds,
	removeCmds,
	restartCmds,
	getCmds,
	loginCmd,
	viewCmds,
	genCmds,
	validateCmd,
	repairCmd,
	chownCmd,
}
