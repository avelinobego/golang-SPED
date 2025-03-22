package schema

import (
	"encoding/xml"
)

// TideEstabLotinfoPerAnt ...
type TideEstabLotinfoPerAnt struct {
	XMLName      xml.Name `xml:"T_ideEstabLot_infoPerAnt"`
	TpInsc       string   `xml:"tpInsc"`
	NrInsc       string   `xml:"nrInsc"`
	CodLotacao   string   `xml:"codLotacao"`
	DetVerbas    []string `xml:"detVerbas"`
	InfoAgNocivo string   `xml:"infoAgNocivo"`
	InfoSimples  string   `xml:"infoSimples"`
}
