package schemas

import (
	"encoding/xml"
)

// Cargo ...
type Cargo struct {
	XMLName  xml.Name `xml:"cargo"`
	DtCargo  string   `xml:"dtCargo"`
	CBOCargo string   `xml:"CBOCargo"`
}
