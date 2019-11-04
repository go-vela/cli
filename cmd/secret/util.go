// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/types/constants"
	"github.com/urfave/cli"
)

// helper function to load global configuration if set via config or environment
func loadGlobal(c *cli.Context) error {

	if len(c.String("engine")) == 0 {
		c.Set("engine", c.GlobalString("secret-engine"))
	}
	if len(c.String("type")) == 0 {
		c.Set("type", c.GlobalString("secret-type"))
	}

	return nil
}

// helper function to validate the user input in the command
func validateCmd(c *cli.Context) error {
	if len(c.String("engine")) == 0 {
		return util.InvalidCommand("engine")
	}
	if len(c.String("type")) == 0 {
		return util.InvalidCommand("type")
	}
	if len(c.String("org")) == 0 {
		return util.InvalidCommand("org")
	}

	return nil
}

// helper function to get the name of the repo, or team for
// sending as the name in the API
func getTypeName(repo, team, stype string) (string, error) {

	if len(repo) == 0 && len(team) == 0 {
		return "", fmt.Errorf("Invalid command: Flag '--repo' or '--team' is not set or is empty")
	}

	// Set name based off user input, If user sets both team and repo
	// default to using the repo value in API request.
	name := ""
	switch stype {
	case constants.SecretShared:
		name = team
	case constants.SecretOrg:
		name = repo
	case constants.SecretRepo:
		name = repo
	}

	return name, nil
}

// helper function to determine if setting value from user input or file
func setValue(s string) (*string, error) {

	if strings.HasPrefix(s, "@") {

		filePath := strings.TrimPrefix(s, "@")
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("Failed to supply valid path: %v", err)
		}

		defer file.Close()
		b, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("Failed to read file: %v", err)
		}

		s = string(b)
	}

	return &s, nil
}
