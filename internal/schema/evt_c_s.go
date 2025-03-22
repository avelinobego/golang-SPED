package schema

import (
	"encoding/xml"
)

// EvtCS ...
type EvtCS struct {
	XMLName       xml.Name `xml:"evtCS"`
	IdAttr        string   `xml:"Id,attr"`
	IdeEvento     string   `xml:"ideEvento"`
	IdeEmpregador string   `xml:"ideEmpregador"`
	InfoCS        *InfoCS  `xml:"infoCS"`
}
