package schemes

import (
	"encoding/xml"
)

// EvtCdBenTerm ...
type EvtCdBenTerm struct {
	XMLName        xml.Name        `xml:"evtCdBenTerm"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      string          `xml:"ideEvento"`
	IdeEmpregador  string          `xml:"ideEmpregador"`
	IdeBeneficio   string          `xml:"ideBeneficio"`
	InfoBenTermino *InfoBenTermino `xml:"infoBenTermino"`
}
