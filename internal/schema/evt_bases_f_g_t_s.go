package schema

import (
	"encoding/xml"
)

// EvtBasesFGTS ...
type EvtBasesFGTS struct {
	XMLName        xml.Name        `xml:"evtBasesFGTS"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      string          `xml:"ideEvento"`
	IdeEmpregador  string          `xml:"ideEmpregador"`
	IdeTrabalhador *IdeTrabalhador `xml:"ideTrabalhador"`
	InfoFGTS       *InfoFGTS       `xml:"infoFGTS"`
}
