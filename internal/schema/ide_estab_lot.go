package schema

import (
	"encoding/xml"
)

// IdeEstabLot ...
type IdeEstabLot struct {
	XMLName     xml.Name     `xml:"ideEstabLot"`
	TpInsc      string       `xml:"tpInsc"`
	NrInsc      string       `xml:"nrInsc"`
	CodLotacao  string       `xml:"codLotacao"`
	DetVerbas   []*DetVerbas `xml:"detVerbas"`
	InfoSimples string       `xml:"infoSimples"`
}
