package schema

import (
	"encoding/xml"
)

// IdeEmpregador ...
type IdeEmpregador struct {
	XMLName xml.Name `xml:"ideEmpregador"`
	TpInsc  string   `xml:"tpInsc"`
	NrInsc  string   `xml:"nrInsc"`
}
