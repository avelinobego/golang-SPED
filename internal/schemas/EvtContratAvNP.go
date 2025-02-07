package schemes

import (
	"encoding/xml"
)

// EvtContratAvNP ...
type EvtContratAvNP struct {
	XMLName       xml.Name     `xml:"evtContratAvNP"`
	IdAttr        string       `xml:"Id,attr"`
	IdeEvento     string       `xml:"ideEvento"`
	IdeEmpregador string       `xml:"ideEmpregador"`
	RemunAvNP     []*RemunAvNP `xml:"remunAvNP"`
}
