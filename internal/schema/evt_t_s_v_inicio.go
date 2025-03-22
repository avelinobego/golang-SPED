package schema

import (
	"encoding/xml"
)

// EvtTSVInicio ...
type EvtTSVInicio struct {
	XMLName       xml.Name       `xml:"evtTSVInicio"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador string         `xml:"ideEmpregador"`
	Trabalhador   *Trabalhador   `xml:"trabalhador"`
	InfoTSVInicio *InfoTSVInicio `xml:"infoTSVInicio"`
}
