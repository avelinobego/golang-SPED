package schema

import (
	"encoding/xml"
)

// Afastamento ...
type Afastamento struct {
	XMLName     xml.Name `xml:"afastamento"`
	DtIniAfast  string   `xml:"dtIniAfast"`
	CodMotAfast string   `xml:"codMotAfast"`
}
