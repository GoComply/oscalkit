// Code generated by https://github.com/GoComply/metaschema; DO NOT EDIT.
package profile

import (
	"encoding/xml"

	"github.com/gocomply/oscalkit/types/oscal/validation_root"

	"github.com/gocomply/oscalkit/types/oscal/nominal_catalog"
)

// Each OSCAL profile is defined by a Profile element
type Profile struct {
	XMLName xml.Name `xml:"http://csrc.nist.gov/ns/oscal/1.0 profile" json:"-"`
	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// Provides information about the publication and availability of the containing document.
	Metadata *Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// An Import element designates a catalog, profile, or other resource to be
	//          included (referenced and potentially modified) by this profile.
	Imports []Import `xml:"import,omitempty" json:"imports,omitempty"`
	// A Merge element merges controls in resolution.
	Merge *Merge `xml:"merge,omitempty" json:"merge,omitempty"`
	// Set parameters or amend controls in resolution
	Modify *Modify `xml:"modify,omitempty" json:"modify,omitempty"`
	// A collection of citations and resource references.
	BackMatter *BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
}

// An Import element designates a catalog, profile, or other resource to be
//          included (referenced and potentially modified) by this profile.
type Import struct {

	// A link to a document or document fragment (actual, nominal or projected)
	Href string `xml:"href,attr,omitempty" json:"href,omitempty"`

	// Specifies which controls to include from the resource (source catalog) being
	//           imported
	Include *Include `xml:"include,omitempty" json:"include,omitempty"`
	// Which controls to exclude from the resource (source catalog) being
	//           imported
	Exclude *Exclude `xml:"exclude,omitempty" json:"exclude,omitempty"`
}

// A Merge element merges controls in resolution.
type Merge struct {

	// A Combine element defines whether and how to combine multiple (competing)
	//         versions of the same control
	Combine *Combine `xml:"combine,omitempty" json:"combine,omitempty"`
	// An As-is element indicates that the controls should be structured in resolution as they are
	//         structured in their source catalogs. It does not contain any elements or attributes.
	AsIs AsIs `xml:"as-is,omitempty" json:"asIs,omitempty"`
	// A Custom element frames a structure for embedding represented controls in resolution.
	Custom *Custom `xml:"custom,omitempty" json:"custom,omitempty"`
}

// A Custom element frames a structure for embedding represented controls in resolution.
type Custom struct {

	// Call a control by its ID
	IdSelectors []Call `xml:"call,omitempty" json:"id-selectors,omitempty"`
	// Select controls by (regular expression) match on ID
	PatternSelectors []Match `xml:"match,omitempty" json:"pattern-selectors,omitempty"`
	// As in catalogs, a group of (selected) controls or of groups of controls
	Groups []Group `xml:"group,omitempty" json:"groups,omitempty"`
}

// As in catalogs, a group of (selected) controls or of groups of controls
type Group struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
	// A partition or component of a control or part
	Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
	// Call a control by its ID
	IdSelectors []Call `xml:"call,omitempty" json:"id-selectors,omitempty"`
	// Select controls by (regular expression) match on ID
	PatternSelectors []Match `xml:"match,omitempty" json:"pattern-selectors,omitempty"`
	// As in catalogs, a group of (selected) controls or of groups of controls
	Groups []Group `xml:"group,omitempty" json:"groups,omitempty"`
}

// Set parameters or amend controls in resolution
type Modify struct {

	// A parameter setting, to be propagated to points of insertion
	ParameterSettings []SetParameter `xml:"set-parameter,omitempty" json:"parameter-settings,omitempty"`
	// An Alter element specifies changes to be made to an included control when a profile is resolved.
	Alterations []Alter `xml:"alter,omitempty" json:"alterations,omitempty"`
}

// Specifies which controls to include from the resource (source catalog) being
//           imported
type Include struct {

	// Include all controls from the imported resource (catalog)
	All *All `xml:"all,omitempty" json:"all,omitempty"`
	// Call a control by its ID
	IdSelectors []Call `xml:"call,omitempty" json:"id-selectors,omitempty"`
	// Select controls by (regular expression) match on ID
	PatternSelectors []Match `xml:"match,omitempty" json:"pattern-selectors,omitempty"`
}

// Which controls to exclude from the resource (source catalog) being
//           imported
type Exclude struct {

	// Call a control by its ID
	IdSelectors []Call `xml:"call,omitempty" json:"id-selectors,omitempty"`
	// Select controls by (regular expression) match on ID
	PatternSelectors []Match `xml:"match,omitempty" json:"pattern-selectors,omitempty"`
}

// A parameter setting, to be propagated to points of insertion
type SetParameter struct {

	// Indicates the value of the 'id' flag on a target parameter; i.e. which parameter to set
	ParamId string `xml:"param-id,attr,omitempty" json:"paramId,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`
	// Another parameter invoking this one
	DependsOn string `xml:"depends-on,attr,omitempty" json:"dependsOn,omitempty"`

	// A placeholder for a missing value, in display.
	Label Label `xml:"label,omitempty" json:"label,omitempty"`
	// Indicates and explains the purpose and use of a parameter
	Descriptions []Usage `xml:"usage,omitempty" json:"descriptions,omitempty"`
	// A formal or informal expression of a constraint or test
	Constraints []Constraint `xml:"constraint,omitempty" json:"constraints,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// A prose statement that provides a recommendation for the use of a parameter.
	Guidance []Guideline `xml:"guideline,omitempty" json:"guidance,omitempty"`
	// Indicates a permissible value for a parameter or property
	Value Value `xml:"value,omitempty" json:"value,omitempty"`
	// Presenting a choice among alternatives
	Select *Select `xml:"select,omitempty" json:"select,omitempty"`
}

// An Alter element specifies changes to be made to an included control when a profile is resolved.
type Alter struct {

	// Value of the 'id' flag on a target control
	ControlId string `xml:"control-id,attr,omitempty" json:"controlId,omitempty"`

	// Specifies elements to be removed from a control, in resolution
	Removals []Remove `xml:"remove,omitempty" json:"removals,omitempty"`
	// Specifies contents to be added into controls, in resolution
	Additions []Add `xml:"add,omitempty" json:"additions,omitempty"`
}

// Specifies contents to be added into controls, in resolution
type Add struct {

	// Where to add the new content with respect to the targeted element (beside it or inside it)
	Position string `xml:"position,attr,omitempty" json:"position,omitempty"`
	// Target location of the addition.
	IdRef string `xml:"id-ref,attr,omitempty" json:"idRef,omitempty"`

	// A title for display and navigation
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// A partition or component of a control or part
	Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
}

// A Combine element defines whether and how to combine multiple (competing)
//         versions of the same control
type Combine struct {
	// How clashing controls should be handled
	Method string `xml:"method,attr,omitempty" json:"method,omitempty"`
	Value  string `xml:",chardata" json:"value,omitempty"`
}

// An As-is element indicates that the controls should be structured in resolution as they are
//         structured in their source catalogs. It does not contain any elements or attributes.

type AsIs string

// Include all controls from the imported resource (catalog)
type All struct {
	// When a control is included, whether its child (dependent) controls are also included.
	WithChildControls string `xml:"with-child-controls,attr,omitempty" json:"withChildControls,omitempty"`
	Value             string `xml:",chardata" json:"value,omitempty"`
}

// Call a control by its ID
type Call struct {
	// Value of the 'id' flag on a target control
	ControlId string `xml:"control-id,attr,omitempty" json:"controlId,omitempty"`

	// When a control is included, whether its child (dependent) controls are also included.
	WithChildControls string `xml:"with-child-controls,attr,omitempty" json:"withChildControls,omitempty"`
	Value             string `xml:",chardata" json:"value,omitempty"`
}

// Select controls by (regular expression) match on ID
type Match struct {
	// A regular expression matching the IDs of one or more controls to be selected
	Pattern string `xml:"pattern,attr,omitempty" json:"pattern,omitempty"`

	// A designation of how a selection of controls in a profile is to be ordered.
	Order string `xml:"order,attr,omitempty" json:"order,omitempty"`

	// When a control is included, whether its child (dependent) controls are also included.
	WithChildControls string `xml:"with-child-controls,attr,omitempty" json:"withChildControls,omitempty"`
	Value             string `xml:",chardata" json:"value,omitempty"`
}

// Specifies elements to be removed from a control, in resolution
type Remove struct {
	// Items to remove, by assigned name
	NameRef string `xml:"name-ref,attr,omitempty" json:"nameRef,omitempty"`

	// Items to remove, by class. A token match.
	ClassRef string `xml:"class-ref,attr,omitempty" json:"classRef,omitempty"`

	// Items to remove, indicated by their IDs
	IdRef string `xml:"id-ref,attr,omitempty" json:"idRef,omitempty"`

	// Items to remove, by the name of the item's type, or generic identifier, e.g.  or
	ItemName string `xml:"item-name,attr,omitempty" json:"itemName,omitempty"`
	Value    string `xml:",chardata" json:"value,omitempty"`
}

type Annotation = validation_root.Annotation

type BackMatter = validation_root.BackMatter

type Constraint = nominal_catalog.Constraint

type Guideline = nominal_catalog.Guideline

type Label = nominal_catalog.Label

type Link = validation_root.Link

type Metadata = validation_root.Metadata

type Param = nominal_catalog.Param

type Part = nominal_catalog.Part

type Prop = validation_root.Prop

type Select = nominal_catalog.Select

type Title = validation_root.Title

type Usage = nominal_catalog.Usage

type Value = nominal_catalog.Value
