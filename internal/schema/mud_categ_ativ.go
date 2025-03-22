package schema

import (
	"encoding/xml"
)

// MudCategAtiv ...
type MudCategAtiv struct {
	XMLName        xml.Name `xml:"mudCategAtiv"`
	CodCateg       string   `xml:"codCateg"`
	NatAtividade   string   `xml:"natAtividade"`
	DtMudCategAtiv string   `xml:"dtMudCategAtiv"`
}
