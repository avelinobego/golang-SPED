package schema

import (
	"encoding/xml"
)

// Endereco ...
type Endereco struct {
	XMLName  xml.Name `xml:"endereco"`
	Brasil   string   `xml:"brasil"`
	Exterior string   `xml:"exterior"`
}
