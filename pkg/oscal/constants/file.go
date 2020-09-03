package constants

// Representing OSCAL file format. XML, JSON, YAML, ...
type DocumentFormat int

const (
	UnknownFormat DocumentFormat = iota
	XmlFormat
	JsonFormat
	YamlFormat
)

type DocumentType int

const (
	UnknownDocument = iota
	CatalogDocument
	ProfileDocument
	SSPDocument
	ComponentDocument
	POAMDocument
)

func (t DocumentType) String() string {
	switch t {
	case CatalogDocument:
		return "Catalog"
	case ProfileDocument:
		return "Profile"
	case SSPDocument:
		return "System Security Plan"
	case ComponentDocument:
		return "Component"
	case POAMDocument:
		return "Plan of Action and Milestones"
	default:
		return "Unrecognized Document Type"
	}
}
