// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipe

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/go-vela/types/pipeline"
	"github.com/go-vela/types/raw"

	"github.com/urfave/cli"
)

// GenCmd defines the command for generating a pipeline.
var GenCmd = cli.Command{
	Name:        "pipe",
	Description: "Use this command to generate a pipeline",
	Usage:       "Generate a vela pipeline in a directory",
	Action:      gen,
	Flags: []cli.Flag{

		// optional flags that can be supplied to a command
		cli.StringFlag{
			Name:  "type,t",
			Usage: "Type of generic pipeline to be generated. (go|node|java)",
		},
		cli.StringFlag{
			Name:  "path,p",
			Usage: "Filename to use to create the secret or secrets",
		},
		cli.BoolFlag{
			Name:  "stages,s",
			Usage: "Define if the pipeline should be generated with stages",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Generate a vela pipeline in your current directory."
		$ {{.HelpName}}
 2. Generate a vela pipeline in a current directory."
		$ {{.HelpName}} -path /path/to/dir/
 3. Generate a node vela pipeline in current directory."
		$ {{.HelpName}} -type node
 4. Generate a java vela pipeline in current directory."
		$ {{.HelpName}} -type java	
 5. Generate a go vela pipeline in a current directory."
		$ {{.HelpName}} -type go
 6. Generate a vela pipeline with stages in current directory."
		$ {{.HelpName}} -stages				
`, cli.CommandHelpTemplate),
}

// helper function to run the init process for a pipeline
func gen(c *cli.Context) error {

	pipe := defaultGoPipe(c.Bool("stages"))
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Unable to get working directory: %v", err)
	}

	if len(c.String("path")) != 0 {

		dir = c.String("path")

		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return fmt.Errorf("Unable to create directory path to config @ %s: %v", dir, err)
		}
	}

	p := fmt.Sprintf("%s/.vela.yml", dir)

	switch c.String("type") {
	case "node":
		pipe = defaultNodePipe(c.Bool("stages"))
	case "java":
		pipe = defaultJavaPipe(c.Bool("stages"))
	}

	data, err := yaml.Marshal(&pipe)
	if err != nil {
		return fmt.Errorf("Unable to create config content: %v", err)
	}

	err = ioutil.WriteFile(p, data, 0600)
	if err != nil {
		return fmt.Errorf("Unable to create yaml config file @ %s: %v", ".vela.yml", err)
	}

	fmt.Printf("\"%s\" %s pipeline generated \n", p, c.String("type"))

	return nil
}

// helper function to define a generic Go pipeline for generation
func defaultGoPipe(s bool) pipeline.Build {

	if s {
		return pipeline.Build{
			Version: "1",
			Stages: pipeline.StageSlice{
				&pipeline.Stage{
					Name: "version",
					Steps: pipeline.ContainerSlice{
						&pipeline.Container{
							Commands: raw.StringSlice{"go version"},
							Image:    "golang:latest",
							Name:     "version",
							Ruleset: pipeline.Ruleset{
								If: pipeline.Rules{Event: []string{"push", "pull_request"}},
							},
						},
					},
				},
			},
		}
	}

	return pipeline.Build{
		Version: "1",
		Steps: pipeline.ContainerSlice{
			&pipeline.Container{
				Commands: raw.StringSlice{"go version"},
				Image:    "golang:latest",
				Name:     "version",
				Ruleset: pipeline.Ruleset{
					If: pipeline.Rules{Event: []string{"push", "pull_request"}},
				},
			},
		},
	}
}

// helper function to define a generic Node pipeline for generation
func defaultNodePipe(s bool) pipeline.Build {

	if s {
		return pipeline.Build{
			Version: "1",
			Stages: pipeline.StageSlice{
				&pipeline.Stage{
					Name: "version",
					Steps: pipeline.ContainerSlice{
						&pipeline.Container{
							Commands: raw.StringSlice{"node --version"},
							Image:    "node:latest",
							Name:     "version",
							Ruleset: pipeline.Ruleset{
								If: pipeline.Rules{Event: []string{"push", "pull_request"}},
							},
						},
					},
				},
			},
		}
	}

	return pipeline.Build{
		Version: "1",
		Steps: pipeline.ContainerSlice{
			&pipeline.Container{
				Commands: raw.StringSlice{"node --version"},
				Image:    "node:latest",
				Name:     "version",
				Ruleset: pipeline.Ruleset{
					If: pipeline.Rules{Event: []string{"push", "pull_request"}},
				},
			},
		},
	}
}

// helper function to define a generic Java pipeline for generation
func defaultJavaPipe(s bool) pipeline.Build {

	if s {
		return pipeline.Build{
			Version: "1",
			Stages: pipeline.StageSlice{
				&pipeline.Stage{
					Name: "version",
					Steps: pipeline.ContainerSlice{
						&pipeline.Container{
							Commands: raw.StringSlice{"java --version"},
							Image:    "openjdk:latest",
							Name:     "version",
							Ruleset: pipeline.Ruleset{
								If: pipeline.Rules{Event: []string{"push", "pull_request"}},
							},
						},
					},
				},
			},
		}
	}

	return pipeline.Build{
		Version: "1",
		Steps: pipeline.ContainerSlice{
			&pipeline.Container{
				Commands: raw.StringSlice{"java --version"},
				Image:    "openjdk:latest",
				Name:     "version",
				Ruleset: pipeline.Ruleset{
					If: pipeline.Rules{Event: []string{"push", "pull_request"}},
				},
			},
		},
	}
}
