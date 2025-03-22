package schema

import (
	"encoding/xml"
)

// Duracao ...
type Duracao struct {
	XMLName   xml.Name `xml:"duracao"`
	TpContr   string   `xml:"tpContr"`
	DtTerm    string   `xml:"dtTerm"`
	ClauAssec string   `xml:"clauAssec"`
	ObjDet    string   `xml:"objDet"`
}
