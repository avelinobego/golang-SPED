package schemes

import (
	"encoding/xml"
)

// TideEstabLot ...
type TideEstabLot struct {
	XMLName      xml.Name `xml:"T_ideEstabLot"`
	TpInsc       string   `xml:"tpInsc"`
	NrInsc       string   `xml:"nrInsc"`
	CodLotacao   string   `xml:"codLotacao"`
	DetVerbas    []string `xml:"detVerbas"`
	InfoAgNocivo string   `xml:"infoAgNocivo"`
	InfoSimples  string   `xml:"infoSimples"`
}
