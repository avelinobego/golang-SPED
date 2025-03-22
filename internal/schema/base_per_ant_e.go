package schema

import (
	"encoding/xml"
)

// BasePerAntE ...
type BasePerAntE struct {
	XMLName   xml.Name `xml:"basePerAntE"`
	TpValorE  string   `xml:"tpValorE"`
	IndIncidE string   `xml:"indIncidE"`
	BaseFGTSE string   `xml:"baseFGTSE"`
	VrFGTSE   string   `xml:"vrFGTSE"`
}
