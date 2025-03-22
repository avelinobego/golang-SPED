package schema

import (
	"encoding/xml"
)

// EvtBaixa ...
type EvtBaixa struct {
	XMLName       xml.Name   `xml:"evtBaixa"`
	IdAttr        string     `xml:"Id,attr"`
	IdeEvento     string     `xml:"ideEvento"`
	IdeEmpregador string     `xml:"ideEmpregador"`
	IdeVinculo    string     `xml:"ideVinculo"`
	InfoBaixa     *InfoBaixa `xml:"infoBaixa"`
}
