package schemas

import (
	"encoding/xml"
)

// IdeVinculo ...
type IdeVinculo struct {
	XMLName   xml.Name `xml:"ideVinculo"`
	CpfTrab   string   `xml:"cpfTrab"`
	Matricula string   `xml:"matricula"`
}
