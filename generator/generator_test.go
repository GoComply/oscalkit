package generator

import (
	"bytes"
	"fmt"
	"net/url"
	"testing"

	"github.com/gocomply/oscalkit/types/oscal/catalog"
	"github.com/gocomply/oscalkit/types/oscal/profile"
)

const (
	temporaryFilePathForCatalogJSON    = "/tmp/catalog.json"
	temporaryFilePathForProfileJSON    = "/tmp/profile.json"
	temporaryFilePathForCatalogsGoFile = "/tmp/catalogs.go"
)

func TestIsHttp(t *testing.T) {

	httpRoute := "http://localhost:3000"
	expectedOutputForHTTP := true

	nonHTTPRoute := "NIST.GOV.JSON"
	expectedOutputForNonHTTP := false

	r, err := url.Parse(httpRoute)
	if err != nil {
		t.Error(err)
	}
	if isHTTPResource(r) != expectedOutputForHTTP {
		t.Error("Invalid output for http routes")
	}

	r, err = url.Parse(nonHTTPRoute)
	if err != nil {
		t.Error(err)
	}
	if isHTTPResource(r) != expectedOutputForNonHTTP {
		t.Error("Invalid output for non http routes")
	}

}

func TestReadInvalidCatalog(t *testing.T) {

	r := bytes.NewReader([]byte(string(`{ "catalog": "some dummy bad json"}`)))
	_, err := ReadCatalog(r)
	if err == nil {
		t.Error("successfully parsed invalid catalog file")
	}
}

func TestGetCatalogInvalidFilePath(t *testing.T) {

	url := "http://[::1]a"
	_, err := GetFilePath(url)
	if err == nil {
		t.Error("should fail")
	}
}

func TestProcessAdditionWithSameClass(t *testing.T) {
	partID := "ac-10_prt"
	class := "guidance"
	alters := []profile.Alter{
		{
			ControlId: "ac-10",
			Additions: []profile.Add{
				profile.Add{
					Parts: []catalog.Part{
						catalog.Part{
							Id:    partID,
							Class: class,
						},
					},
				},
			},
		},
		profile.Alter{
			ControlId: "ac-10.1",
			Additions: []profile.Add{
				profile.Add{
					Parts: []catalog.Part{
						catalog.Part{
							Id:    partID,
							Class: class,
						},
					},
				},
			},
		},
	}
	c := catalog.Catalog{
		Groups: []catalog.Group{
			catalog.Group{
				Controls: []catalog.Control{
					catalog.Control{
						Id: "ac-10",
						Parts: []catalog.Part{
							catalog.Part{
								Id:    partID,
								Class: class,
							},
						},
						Controls: []catalog.Control{
							catalog.Control{
								Id: "ac-10.1",
								Parts: []catalog.Part{
									catalog.Part{
										Id:    partID,
										Class: class,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	o := ProcessAlterations(alters, &c)
	for _, g := range o.Groups {
		for _, c := range g.Controls {
			for i := range c.Parts {
				expected := fmt.Sprintf("%s_%d", partID, i+1)
				if c.Parts[i].Id != expected {
					t.Errorf("%s and %s are not identical", c.Parts[i].Id, expected)
					return
				}
			}
			for i, sc := range c.Controls {
				expected := fmt.Sprintf("%s_%d", partID, i+1)
				if sc.Parts[i].Id != expected {
					t.Errorf("%s and %s are not identical", sc.Parts[i].Id, expected)
					return
				}
			}
		}
	}
}

func TestProcessAdditionWithDifferentPartClass(t *testing.T) {

	ctrlID := "ac-10"
	subctrlID := "ac-10.1"
	partID := "ac-10_stmt.a"

	alters := []profile.Alter{
		profile.Alter{
			ControlId: ctrlID,
			Additions: []profile.Add{
				profile.Add{
					Parts: []catalog.Part{
						catalog.Part{
							Id:    partID,
							Class: "c1",
						},
					},
				},
			},
		},
		profile.Alter{
			ControlId: subctrlID,
			Additions: []profile.Add{
				profile.Add{
					Parts: []catalog.Part{
						catalog.Part{
							Id:    partID,
							Class: "c2",
						},
					},
				},
			},
		},
	}
	c := catalog.Catalog{
		Groups: []catalog.Group{
			catalog.Group{
				Controls: []catalog.Control{
					catalog.Control{
						Id: ctrlID,
						Parts: []catalog.Part{
							catalog.Part{
								Id:    partID,
								Class: "c3",
							},
						},
						Controls: []catalog.Control{
							catalog.Control{
								Id: subctrlID,
								Parts: []catalog.Part{
									catalog.Part{
										Id:    partID,
										Class: "c4",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	o := ProcessAlterations(alters, &c)
	if len(o.Groups[0].Controls[0].Parts) != 2 {
		t.Error("parts for controls not getting added properly")
	}
	if len(o.Groups[0].Controls[0].Controls[0].Parts) != 2 {
		t.Error("parts for sub-controls not getting added properly")
	}

}

func failTest(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
