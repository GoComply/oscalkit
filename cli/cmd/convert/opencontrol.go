package convert

import (
	"fmt"

	"github.com/gocomply/oscalkit/pkg/opencontrol"
	"github.com/gocomply/oscalkit/pkg/oscal/constants"
	"github.com/gocomply/oscalkit/pkg/oscal_source"
	"github.com/urfave/cli"
)

var ConvertOpenControl = cli.Command{
	Name:        "opencontrol",
	Usage:       "convert OSCAL Catalog file to OpenControl Standard file",
	Description: `Convert OSCAL-formatted Catalog file to OpenControl-formatted standard file.`,
	ArgsUsage:   "[input-catalog-file]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "output-file, o",
			Usage:       "Output path for converted file",
			Destination: &outputPath,
		},
	},
	Before: func(c *cli.Context) error {
		if c.NArg() != 1 {
			return cli.NewExitError("gocomply_oscalkit convert opencontrol requires exactly one argument", 1)
		}

		if outputPath == "" {
			return cli.NewExitError("Please provide -o output destination", 1)
		}

		return nil
	},
	Action: func(c *cli.Context) error {
		for _, catalogPath := range c.Args() {
			source, err := oscal_source.Open(catalogPath)
			if err != nil {
				return cli.NewExitError(fmt.Sprintf("Could not parse input file: %s", err), 1)
			}
			o := source.OSCAL()
			if o.DocumentType() != constants.CatalogDocument {
				return cli.NewExitError(fmt.Sprintf("Unexpected OSCAL %s found, expected Catalog document", o.DocumentType().String()), 1)
			}
			standard, err := opencontrol.NewStandard(o.Catalog)
			if err != nil {
				return cli.NewExitError(fmt.Sprintf("Cannot convert to opencontrol standard: %v", err), 1)
			}
			return standard.SaveToFile(outputPath)
		}
		return nil
	},
}
