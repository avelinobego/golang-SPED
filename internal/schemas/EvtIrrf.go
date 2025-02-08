package schemas

import (
	"encoding/xml"
)

// EvtIrrf ...
type EvtIrrf struct {
	XMLName       xml.Name  `xml:"evtIrrf"`
	IdAttr        string    `xml:"Id,attr"`
	IdeEvento     string    `xml:"ideEvento"`
	IdeEmpregador string    `xml:"ideEmpregador"`
	InfoIRRF      *InfoIRRF `xml:"infoIRRF"`
}
