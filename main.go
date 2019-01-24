package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// App constants for the CLI
const (
	AppName    = "gostitch"
	AppUsage   = "stitch files within various directories"
	AppVersion = "1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = AppVersion
	app.Action = func(c *cli.Context) error {
		fmt.Println(`type "gostitch help" for extra information`)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "update",
			Usage: "creates/updates YAML configuration file for file stitching",
			Action: func(c *cli.Context) error {
				return StitchInit()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
