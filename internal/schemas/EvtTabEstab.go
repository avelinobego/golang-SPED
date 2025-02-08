package schemas

import (
	"encoding/xml"
)

// EvtTabEstab ...
type EvtTabEstab struct {
	XMLName       xml.Name   `xml:"evtTabEstab"`
	IdAttr        string     `xml:"Id,attr"`
	IdeEvento     string     `xml:"ideEvento"`
	IdeEmpregador string     `xml:"ideEmpregador"`
	InfoEstab     *InfoEstab `xml:"infoEstab"`
}
