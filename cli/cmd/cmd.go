package cmd

import (
	"fmt"
	"os"

	"github.com/gocomply/oscalkit/cli/cmd/convert"
	"github.com/gocomply/oscalkit/cli/cmd/generate"
	"github.com/gocomply/oscalkit/cli/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Execute ...
func Execute() error {
	appVersion := fmt.Sprintf("%s (Built: %s)\n", version.Version, version.Date)

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(appVersion)
	}

	app := cli.NewApp()
	app.Name = "gocomply_oscalkit"
	app.Version = appVersion
	app.Usage = "OSCAL toolkit"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable debug command output",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}
	app.Commands = []cli.Command{
		Info,
		convert.Convert,
		Diff,
		Validate,
		Sign,
		generate.Generate,
	}

	return app.Run(os.Args)
}
