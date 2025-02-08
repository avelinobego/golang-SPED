package schemas

import (
	"encoding/xml"
)

// EvtCdBenIn ...
type EvtCdBenIn struct {
	XMLName       xml.Name       `xml:"evtCdBenIn"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador string         `xml:"ideEmpregador"`
	Beneficiario  *Beneficiario  `xml:"beneficiario"`
	InfoBenInicio *InfoBenInicio `xml:"infoBenInicio"`
}
