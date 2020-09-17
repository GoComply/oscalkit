package oscal_diff

import (
	"fmt"

	"github.com/pmezard/go-difflib/difflib"
	"github.com/gocomply/oscalkit/types/oscal"
	"github.com/gocomply/oscalkit/pkg/oscal/constants"
	"github.com/davecgh/go-spew/spew"
)

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
}

func Diff(a *oscal.OSCAL, b *oscal.OSCAL) (string, error) {
	if a.DocumentType() != b.DocumentType() {
		return "", fmt.Errorf("Could not compare OSCAL resources, type mismatch '%s' vs '%s'", a.DocumentType(), b.DocumentType())
	}

	switch a.DocumentType() {
	case constants.CatalogDocument:
		if a.Catalog.XMLName.Space != b.Catalog.XMLName.Space {
			if a.Catalog.XMLName.Space == "" {
				a.Catalog.XMLName = b.Catalog.XMLName
			}
			if b.Catalog.XMLName.Space == "" {
				b.Catalog.XMLName = a.Catalog.XMLName
			}
		}
	}

	as := spewConfig.Sdump(a)
	bs := spewConfig.Sdump(b)

	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(as),
		B:        difflib.SplitLines(bs),
		FromFile: "a1",
		ToFile:   "a2",
		Context:  3,
	}
	return difflib.GetUnifiedDiffString(diff)
}
