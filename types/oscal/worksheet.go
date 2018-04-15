// oscalkit - OSCAL conversion utility
// Written in 2017 by Andrew Weiss <andrew.weiss@docker.com>

// To the extent possible under law, the author(s) have dedicated all copyright
// and related and neighboring rights to this software to the public domain worldwide.
// This software is distributed without any warranty.

// You should have received a copy of the CC0 Public Domain Dedication along with this software.
// If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.

package oscal

import (
	"encoding/json"
	"encoding/xml"

	yaml "gopkg.in/yaml.v2"
)

// Worksheet ...
// Create type alias and override component method
type Worksheet struct {
	XMLName         xml.Name               `xml:"http://csrc.nist.gov/ns/oscal/1.0 worksheet" json:"-" yaml:"-"`
	OptionalClass   string                 `xml:"class,attr,omitempty" json:"class,omitempty" yaml:"class,omitempty"`
	ID              string                 `xml:"id,attr,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Title           *Raw                   `xml:"title" json:"title" yaml:"title"`
	Declarations    *FrameworkDeclarations `xml:"declarations,omitempty" json:"declarations,omitempty" yaml:"declarations,omitempty"`
	Sections        []Section              `xml:"section,omitempty" json:"sections,omitempty" yaml:"sections,omitempty"`
	ComponentGroups []ComponentGroup       `xml:"group,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Components      []Components           `xml:"component,omitempty" json:"components,omitempty" yaml:"components,omitempty"`
	References      *References            `xml:"references,omitempty" json:"references,omitempty" yaml:"references,omitempty"`
}

// Component ...
func (Worksheet) Component() string {
	return "worksheet"
}

// RawXML ...
func (w *Worksheet) RawXML(prettify bool) ([]byte, error) {
	if prettify {
		return xml.MarshalIndent(w, "", "  ")
	}
	return xml.Marshal(w)
}

// RawJSON ...
func (w *Worksheet) RawJSON(prettify bool) ([]byte, error) {
	if prettify {
		return json.MarshalIndent(w, "", "  ")
	}
	return json.Marshal(w)
}

// RawYAML ...
func (w *Worksheet) RawYAML() ([]byte, error) {
	return yaml.Marshal(w)
}