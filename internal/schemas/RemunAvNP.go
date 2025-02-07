package schemes

import (
	"encoding/xml"
)

// RemunAvNP ...
type RemunAvNP struct {
	XMLName    xml.Name `xml:"remunAvNP"`
	TpInsc     string   `xml:"tpInsc"`
	NrInsc     string   `xml:"nrInsc"`
	CodLotacao string   `xml:"codLotacao"`
	VrBcCp00   string   `xml:"vrBcCp00"`
	VrBcCp15   string   `xml:"vrBcCp15"`
	VrBcCp20   string   `xml:"vrBcCp20"`
	VrBcCp25   string   `xml:"vrBcCp25"`
	VrBcCp13   string   `xml:"vrBcCp13"`
	VrBcFgts   string   `xml:"vrBcFgts"`
	VrDescCP   string   `xml:"vrDescCP"`
}
