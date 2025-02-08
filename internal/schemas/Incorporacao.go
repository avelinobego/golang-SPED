package schemas

import (
	"encoding/xml"
)

// Incorporacao ...
type Incorporacao struct {
	XMLName   xml.Name `xml:"incorporacao"`
	TpInsc    string   `xml:"tpInsc"`
	NrInsc    string   `xml:"nrInsc"`
	MatIncorp string   `xml:"matIncorp"`
}
