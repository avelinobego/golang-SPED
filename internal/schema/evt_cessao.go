package schema

import (
	"encoding/xml"
)

// EvtCessao ...
type EvtCessao struct {
	XMLName       xml.Name    `xml:"evtCessao"`
	IdAttr        string      `xml:"Id,attr"`
	IdeEvento     string      `xml:"ideEvento"`
	IdeEmpregador string      `xml:"ideEmpregador"`
	IdeVinculo    string      `xml:"ideVinculo"`
	InfoCessao    *InfoCessao `xml:"infoCessao"`
}
