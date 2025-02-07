package schemes

import (
	"encoding/xml"
)

// DetInfoPerRef ...
type DetInfoPerRef struct {
	XMLName    xml.Name `xml:"detInfoPerRef"`
	Ind13      string   `xml:"ind13"`
	TpVrPerRef int8     `xml:"tpVrPerRef"`
	VrPerRef   string   `xml:"vrPerRef"`
}
