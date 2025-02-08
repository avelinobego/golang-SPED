package schemas

import (
	"encoding/xml"
)

// Emitente ...
type Emitente struct {
	XMLName xml.Name `xml:"emitente"`
	NmEmit  string   `xml:"nmEmit"`
	IdeOC   int8     `xml:"ideOC"`
	NrOC    string   `xml:"nrOC"`
	UfOC    string   `xml:"ufOC"`
}
