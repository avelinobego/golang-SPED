package schemes

import (
	"encoding/xml"
)

// EvtCdBenAlt ...
type EvtCdBenAlt struct {
	XMLName          xml.Name          `xml:"evtCdBenAlt"`
	IdAttr           string            `xml:"Id,attr"`
	IdeEvento        string            `xml:"ideEvento"`
	IdeEmpregador    string            `xml:"ideEmpregador"`
	IdeBeneficio     string            `xml:"ideBeneficio"`
	InfoBenAlteracao *InfoBenAlteracao `xml:"infoBenAlteracao"`
}
