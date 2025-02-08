package schemas

import (
	"encoding/xml"
)

// EvtExpRisco ...
type EvtExpRisco struct {
	XMLName       xml.Name      `xml:"evtExpRisco"`
	IdAttr        string        `xml:"Id,attr"`
	IdeEvento     string        `xml:"ideEvento"`
	IdeEmpregador string        `xml:"ideEmpregador"`
	IdeVinculo    string        `xml:"ideVinculo"`
	InfoExpRisco  *InfoExpRisco `xml:"infoExpRisco"`
}
