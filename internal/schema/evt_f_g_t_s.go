package schema

import (
	"encoding/xml"
)

// EvtFGTS ...
type EvtFGTS struct {
	XMLName       xml.Name  `xml:"evtFGTS"`
	IdAttr        string    `xml:"Id,attr"`
	IdeEvento     string    `xml:"ideEvento"`
	IdeEmpregador string    `xml:"ideEmpregador"`
	InfoFGTS      *InfoFGTS `xml:"infoFGTS"`
}
