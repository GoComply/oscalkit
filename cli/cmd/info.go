package cmd

import (
	"errors"
	"fmt"

	"github.com/gocomply/oscalkit/pkg/oscal/constants"
	"github.com/gocomply/oscalkit/pkg/oscal_source"
	"github.com/gocomply/oscalkit/types/oscal"
	"github.com/gocomply/oscalkit/types/oscal/catalog"
	"github.com/gocomply/oscalkit/types/oscal/profile"
	"github.com/urfave/cli"
)

var Info = cli.Command{
	Name:      "info",
	Usage:     "Provides information about particular OSCAL resource",
	ArgsUsage: "file [file...]",
	Before: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return cli.NewExitError("No file provided", 1)
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		for _, filePath := range c.Args() {
			os, err := oscal_source.Open(filePath)
			if err != nil {
				return cli.NewExitError(fmt.Sprintf("Could not open oscal file: %v", err), 1)
			}
			defer os.Close()

			err = printInfo(os.OSCAL())
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func printInfo(o *oscal.OSCAL) error {
	switch o.DocumentType() {
	case constants.SSPDocument:
		fmt.Println("OSCAL System Security Plan")
		fmt.Println("UUID:\t", o.SystemSecurityPlan.Uuid)
		printMetadata(o.SystemSecurityPlan.Metadata)
		return nil
	case constants.ComponentDocument:
		fmt.Println("OSCAL Component (represents information about particular software asset/component)")
		printMetadata(o.Component.Metadata)
		return nil
	case constants.ProfileDocument:
		fmt.Println("OSCAL Profile (represents tailoring of controls from OSCAL catalog(s) or profile(s))")
		fmt.Println("UUID:\t", o.Profile.Uuid)
		printMetadata(o.Profile.Metadata)
		return printImports(o.Profile)
	case constants.CatalogDocument:
		fmt.Println("OSCAL Catalog (represents library of control assessment objectives and activities)")
		fmt.Println("UUID:\t", o.Catalog.Uuid)
		printMetadata(o.Catalog.Metadata)
		return nil
	case constants.POAMDocument:
		fmt.Printf("OSCAL %s (represents the known risks for a specific system, as well as the identified deviations, remediation plan, and disposition status of each risk)\n", o.DocumentType().String())
		fmt.Println("UUID:\t", o.PlanOfActionAndMilestones.Uuid)
		printMetadata(o.PlanOfActionAndMilestones.Metadata)
		return nil
	case constants.AssessmentPlanDocument:
		fmt.Printf("OSCAL %s (represents the planning of a periodic or continuous assessment)\n", o.DocumentType().String())
		fmt.Println("UUID:\t", o.AssessmentPlan.Uuid)
		printMetadata(o.AssessmentPlan.Metadata)
		return nil
	case constants.AssessmentResultsDocument:
		fmt.Printf("OSCAL %s (represents the findings of a periodic or continuous assessment of a specific system)\n", o.DocumentType().String())
		fmt.Println("UUID:\t", o.AssessmentResults.Uuid)
		printMetadata(o.AssessmentResults.Metadata)
		return nil
	}
	return errors.New("Unrecognized OSCAL resource")
}

func printMetadata(m *catalog.Metadata) {
	if m == nil {
		return
	}
	fmt.Println("Metadata:")
	fmt.Println("\tTitle:\t\t\t", m.Title.Raw)
	if m.Published != "" {
		fmt.Println("\tPublished:\t\t", m.Published)
	}
	if m.LastModified != "" {
		fmt.Println("\tLast Modified:\t\t", m.LastModified)
	}
	if m.Version != "" {
		fmt.Println("\tDocument Version:\t", m.Version)
	}
	if m.OscalVersion != "" {
		fmt.Println("\tOSCAL Version:\t\t", m.OscalVersion)
	}
}

func printImports(p *profile.Profile) error {
	fmt.Println("This profile builds on top of (Imports):")
	for i, imp := range p.Imports {
		fmt.Printf("  (%d)", i+1)
		if imp.IsDocumentFragment() {
			resource, err := p.GetDocumentFragment(imp.Href)
			if err != nil {
				return err
			}
			if resource == nil {
				return fmt.Errorf("Could not resolve profile import %s", imp.Href)
			}

			if resource.Title != nil {
				fmt.Println("\tTitle:\t\t\t", resource.Title.Raw)
			}
			if resource.Desc != "" {
				fmt.Println("\tDesc:\t\t\t", resource.Desc)
			}
			for _, rlink := range resource.Rlinks {
				fmt.Println("\tLink:\t\t\t", rlink.Href)
			}

		} else {
			fmt.Printf("\thref=%s\n", imp.Href)
		}
	}
	return nil
}
