package schemes

import (
	"encoding/xml"
)

// EvtReintegr ...
type EvtReintegr struct {
	XMLName       xml.Name      `xml:"evtReintegr"`
	IdAttr        string        `xml:"Id,attr"`
	IdeEvento     string        `xml:"ideEvento"`
	IdeEmpregador string        `xml:"ideEmpregador"`
	IdeVinculo    string        `xml:"ideVinculo"`
	InfoReintegr  *InfoReintegr `xml:"infoReintegr"`
}
