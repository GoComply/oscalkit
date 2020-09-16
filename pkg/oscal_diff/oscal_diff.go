package oscal_diff

import (
	"github.com/pmezard/go-difflib/difflib"
	"github.com/gocomply/oscalkit/types/oscal"
	"github.com/davecgh/go-spew/spew"
)

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
}

func Diff(a *oscal.OSCAL, b *oscal.OSCAL) (string, error) {
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
