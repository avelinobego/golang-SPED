package schemes

import (
	"encoding/xml"
)

// Termino ...
type Termino struct {
	XMLName xml.Name `xml:"termino"`
	DtTerm  string   `xml:"dtTerm"`
}
