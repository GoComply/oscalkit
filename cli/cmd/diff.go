package cmd

import (
	"fmt"

	"github.com/gocomply/oscalkit/pkg/oscal_diff"
	"github.com/gocomply/oscalkit/pkg/oscal_source"
	"github.com/urfave/cli"
)

var Diff = cli.Command{
	Name:      "diff",
	Usage:     "Compare OSCAL resources semantically item by item",
	ArgsUsage: "FILE1 FILE2",
	Before: func(c *cli.Context) error {
		if c.NArg() != 2 {
			return cli.NewExitError("Please supply exactly file paths as command-line arguments", 1)
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		fileA, fileB := c.Args()[0], c.Args()[1]
		osA, err := oscal_source.Open(fileA)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Could not open oscal file %s: %v", fileA, err), 1)
		}
		defer osA.Close()

		osB, err := oscal_source.Open(fileB)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Could not open oscal file %s: %v", fileB, err), 1)
		}
		defer osB.Close()

		text, err := oscal_diff.Diff(osA.OSCAL(), osB.OSCAL())
		if err != nil {
			return err
		}
		if text != "" {
			return cli.NewExitError(text, 2)
		}
		return nil
	},
}
