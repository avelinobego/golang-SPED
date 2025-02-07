package schemes

import (
	"encoding/xml"
)

// IdeLotacao ...
type IdeLotacao struct {
	XMLName      xml.Name      `xml:"ideLotacao"`
	CodLotacao   string        `xml:"codLotacao"`
	TpLotacao    string        `xml:"tpLotacao"`
	TpInsc       string        `xml:"tpInsc"`
	NrInsc       string        `xml:"nrInsc"`
	InfoBaseFGTS *InfoBaseFGTS `xml:"infoBaseFGTS"`
}
