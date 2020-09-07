// Code generated by https://github.com/GoComply/metaschema; DO NOT EDIT.
package component_definition

import (
	"github.com/gocomply/oscalkit/types/oscal/validation_root"

	"github.com/gocomply/oscalkit/types/oscal/validation_common_root"
)

// A collection of component descriptions, which may optionally be grouped by capability.
type ComponentDefinition struct {

	// Loads a component definition from another resource.
	ImportComponentDefinitions []ImportComponentDefinition `xml:"import-component-definition,omitempty" json:"import-component-definitions,omitempty"`
	// Provides information about the publication and availability of the containing document.
	Metadata *Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// A defined component that can be part of an implemented system.
	Components []Component `xml:"component,omitempty" json:"components,omitempty"`
	// A grouping of other components and/or capabilities.
	Capabilities []Capability `xml:"capability,omitempty" json:"capabilities,omitempty"`
	// A collection of citations and resource references.
	BackMatter *BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
}

// A defined component that can be part of an implemented system.
type Component struct {

	// A unique identifier for a component.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
	// The component's short, human-readable name.
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
	// A category describing the purpose of the component.
	ComponentType string `xml:"component-type,attr,omitempty" json:"componentType,omitempty"`

	// A longer name for the component.
	Title Title `xml:"title,omitempty" json:"title,omitempty"`
	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// Defines a role associated with a party or parties that has responsibility for the component.
	ResponsibleParties []ResponsibleParty `xml:"responsible-party,omitempty" json:"responsible-parties,omitempty"`
	// Defines how the component or capability supports a set of controls.
	ControlImplementations []ControlImplementation `xml:"control-implementation,omitempty" json:"control-implementations,omitempty"`
}

// A grouping of other components and/or capabilities.
type Capability struct {

	// A unique identifier for a capability.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
	// The capability's human-readable name.
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`

	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// TBD
	IncorporatesComponents []IncorporatesComponent `xml:"incorporates-component,omitempty" json:"incorporates-components,omitempty"`
	// Defines how the component or capability supports a set of controls.
	ControlImplementations []ControlImplementation `xml:"control-implementation,omitempty" json:"control-implementations,omitempty"`
}

// Defines how the component or capability supports a set of controls.
type ControlImplementation struct {

	// A unique identifier for the set of implemented controls.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
	// A URL reference to the source catalog or profile for which this component is implementing controls for.
	Source string `xml:"source,attr,omitempty" json:"source,omitempty"`

	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// Describes how the component implements an individual control.
	ImplementedRequirements []ImplementedRequirement `xml:"implemented-requirement,omitempty" json:"implemented-requirements,omitempty"`
}

// Describes how the component implements an individual control.
type ImplementedRequirement struct {

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
	// A reference to a control identifier.
	ControlId string `xml:"control-id,attr,omitempty" json:"controlId,omitempty"`

	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// A reference to one or more roles with responsibility for performing a function relative to the control.
	ResponsibleRoles []ResponsibleRole `xml:"responsible-role,omitempty" json:"responsible-roles,omitempty"`
	// Identifies the parameter that will be filled in by the enclosed value element.
	SetParameters []SetParameter `xml:"set-parameter,omitempty" json:"set-parameters,omitempty"`
	// Identifies which statements within a control are addressed.
	Statements []Statement `xml:"statement,omitempty" json:"statements,omitempty"`
}

// Identifies which statements within a control are addressed.
type Statement struct {

	// A reference to the specific implemented statement associated with a control.
	StatementId string `xml:"statement-id,attr,omitempty" json:"statementId,omitempty"`
	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// A description supporting the parent item.
	Description *Description `xml:"description,omitempty" json:"description,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Links []Link `xml:"link,omitempty" json:"links,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
	// A reference to one or more roles with responsibility for performing a function relative to the control.
	ResponsibleRoles []ResponsibleRole `xml:"responsible-role,omitempty" json:"responsible-roles,omitempty"`
}

// Loads a component definition from another resource.
type ImportComponentDefinition struct {
	// A link to a resource that defines a set of components and/or capabilities to import into this collection.
	Href  string `xml:"href,attr,omitempty" json:"href,omitempty"`
	Value string `xml:",chardata" json:"value,omitempty"`
}

type Annotation = validation_root.Annotation

type BackMatter = validation_root.BackMatter

type Description = validation_root.Description

type IncorporatesComponent = validation_common_root.IncorporatesComponent

type Link = validation_root.Link

type Metadata = validation_root.Metadata

type Prop = validation_root.Prop

type Remarks = validation_root.Remarks

type ResponsibleParty = validation_root.ResponsibleParty

type ResponsibleRole = validation_common_root.ResponsibleRole

type SetParameter = validation_common_root.SetParameter

type Title = validation_root.Title
