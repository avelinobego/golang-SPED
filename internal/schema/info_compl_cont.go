package schema

import (
	"encoding/xml"
)

// InfoComplCont ...
type InfoComplCont struct {
	XMLName      xml.Name `xml:"infoComplCont"`
	CodCBO       string   `xml:"codCBO"`
	NatAtividade string   `xml:"natAtividade"`
	QtdDiasTrab  string   `xml:"qtdDiasTrab"`
}
