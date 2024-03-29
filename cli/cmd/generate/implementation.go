package generate

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocomply/oscalkit/generator"
	"github.com/gocomply/oscalkit/impl"
	"github.com/gocomply/oscalkit/templates"
	"github.com/gocomply/oscalkit/types/oscal/implementation"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var profile string
var excelSheet string

// Implementation generates implemntation
var Implementation = cli.Command{
	Name:  "implementation",
	Usage: "generates go code for implementation against provided profile and excel sheet",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "excel, e",
			Usage:       "excel sheet to get component configs",
			Destination: &excelSheet,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "output filename",
			Destination: &outputFileName,
			Value:       "implementation.go",
		},
		cli.StringFlag{
			Name:        "package, pkg",
			Usage:       "package name for generated go file (default is oscalkit)",
			Destination: &packageName,
			Value:       "oscalkit",
		},
	},
	Before: func(c *cli.Context) error {
		if excelSheet == "" {
			return cli.NewExitError("gocomply_oscalkit implementation is missing --excel flag", 1)
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		err := validatePackageName(packageName)
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		excelF, err := generator.GetFilePath(excelSheet)
		if err != nil {
			return err
		}
		b, err := ioutil.ReadFile(excelF)
		if err != nil {
			return err
		}

		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return fmt.Errorf("cannot create file for implementation: err: %v", err)
		}
		defer outputFile.Close()

		var records [][]string
		reader := bytes.NewReader(b)
		r := csv.NewReader(reader)
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)
		}

		catalog := impl.NISTCatalog{ID: "NIST_SP-800-53"}
		implementationData := impl.GenerateImplementation(records, &catalog)
		t, err := templates.GetImplementationTemplate()
		if err != nil {
			return fmt.Errorf("cannot get implementation template err %v", err)
		}
		err = t.Execute(outputFile, struct {
			Implementation implementation.Implementation
			PackageName    string
		}{implementationData, packageName})
		if err != nil {
			return err
		}
		b, err = ioutil.ReadFile(outputFileName)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("cannot open %s file", outputFileName), 1)
		}
		b, err = format.Source(b)
		if err != nil {
			logrus.Warn(fmt.Sprintf("cannot format %s file", outputFileName))
			return cli.NewExitError(err, 1)
		}
		err = ioutil.WriteFile(outputFileName, b, 0)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("cannot write formmated "), 1)
		}
		return nil
	},
}
