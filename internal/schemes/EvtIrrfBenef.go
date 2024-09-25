package schemes

import (
	"encoding/xml"
)

// EvtIrrfBenef ...
type EvtIrrfBenef struct {
	XMLName        xml.Name        `xml:"evtIrrfBenef"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      *IdeEvento      `xml:"ideEvento"`
	IdeEmpregador  string          `xml:"ideEmpregador"`
	IdeTrabalhador *IdeTrabalhador `xml:"ideTrabalhador"`
}
