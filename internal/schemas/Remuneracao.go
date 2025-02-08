package schemas

import (
	"encoding/xml"
)

// Remuneracao ...
type Remuneracao struct {
	XMLName    xml.Name `xml:"remuneracao"`
	DtRemun    string   `xml:"dtRemun"`
	VrSalFx    string   `xml:"vrSalFx"`
	UndSalFixo string   `xml:"undSalFixo"`
	DscSalVar  string   `xml:"dscSalVar"`
}
