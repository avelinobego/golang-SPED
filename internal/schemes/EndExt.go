package schemes

import (
	"encoding/xml"
)

// EndExt ...
type EndExt struct {
	XMLName      xml.Name `xml:"endExt"`
	EndDscLograd string   `xml:"endDscLograd"`
	EndNrLograd  string   `xml:"endNrLograd"`
	EndComplem   string   `xml:"endComplem"`
	EndBairro    string   `xml:"endBairro"`
	EndCidade    string   `xml:"endCidade"`
	EndEstado    string   `xml:"endEstado"`
	EndCodPostal string   `xml:"endCodPostal"`
	Telef        string   `xml:"telef"`
}
