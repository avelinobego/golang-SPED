package schemes

import (
	"encoding/xml"
)

// InfoAmb ...
type InfoAmb struct {
	XMLName  xml.Name `xml:"infoAmb"`
	LocalAmb int8     `xml:"localAmb"`
	DscSetor string   `xml:"dscSetor"`
	TpInsc   string   `xml:"tpInsc"`
	NrInsc   string   `xml:"nrInsc"`
}
