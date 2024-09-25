package schemes

import (
	"encoding/xml"
)

// EvtDeslig ...
type EvtDeslig struct {
	XMLName       xml.Name    `xml:"evtDeslig"`
	IdAttr        string      `xml:"Id,attr"`
	IdeEvento     string      `xml:"ideEvento"`
	IdeEmpregador string      `xml:"ideEmpregador"`
	IdeVinculo    string      `xml:"ideVinculo"`
	InfoDeslig    *InfoDeslig `xml:"infoDeslig"`
}
