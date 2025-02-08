package schemas

import (
	"encoding/xml"
)

// EvtTSVAltContr ...
type EvtTSVAltContr struct {
	XMLName           xml.Name          `xml:"evtTSVAltContr"`
	IdAttr            string            `xml:"Id,attr"`
	IdeEvento         string            `xml:"ideEvento"`
	IdeEmpregador     string            `xml:"ideEmpregador"`
	IdeTrabSemVinculo string            `xml:"ideTrabSemVinculo"`
	InfoTSVAlteracao  *InfoTSVAlteracao `xml:"infoTSVAlteracao"`
}
