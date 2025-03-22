package schema

import (
	"encoding/xml"
)

// InfoCompl ...
type InfoCompl struct {
	XMLName      xml.Name       `xml:"infoCompl"`
	CodCBO       string         `xml:"codCBO"`
	NatAtividade string         `xml:"natAtividade"`
	Remuneracao  []*Remuneracao `xml:"remuneracao"`
	InfoVinc     *InfoVinc      `xml:"infoVinc"`
	InfoTerm     *InfoTerm      `xml:"infoTerm"`
}
