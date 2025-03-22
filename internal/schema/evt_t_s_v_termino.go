package schema

import (
	"encoding/xml"
)

// EvtTSVTermino ...
type EvtTSVTermino struct {
	XMLName           xml.Name        `xml:"evtTSVTermino"`
	IdAttr            string          `xml:"Id,attr"`
	IdeEvento         string          `xml:"ideEvento"`
	IdeEmpregador     string          `xml:"ideEmpregador"`
	IdeTrabSemVinculo string          `xml:"ideTrabSemVinculo"`
	InfoTSVTermino    *InfoTSVTermino `xml:"infoTSVTermino"`
}
