package schema

import (
	"encoding/xml"
)

// IdeEstab ...
type IdeEstab struct {
	XMLName      xml.Name `xml:"ideEstab"`
	TpInsc       string   `xml:"tpInsc"`
	NrInsc       string   `xml:"nrInsc"`
	RemunPerApur []string `xml:"remunPerApur"`
}
