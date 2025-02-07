package schemes

import (
	"encoding/xml"
)

// EvtCdBenefIn ...
type EvtCdBenefIn struct {
	XMLName       xml.Name       `xml:"evtCdBenefIn"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador *IdeEmpregador `xml:"ideEmpregador"`
	Beneficiario  *Beneficiario  `xml:"beneficiario"`
}
