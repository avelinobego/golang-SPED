package schemas

import (
	"encoding/xml"
)

// EvtReativBen ...
type EvtReativBen struct {
	XMLName       xml.Name    `xml:"evtReativBen"`
	IdAttr        string      `xml:"Id,attr"`
	IdeEvento     string      `xml:"ideEvento"`
	IdeEmpregador string      `xml:"ideEmpregador"`
	IdeBeneficio  string      `xml:"ideBeneficio"`
	InfoReativ    *InfoReativ `xml:"infoReativ"`
}
