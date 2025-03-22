package schema

import (
	"encoding/xml"
)

// EvtExclusao ...
type EvtExclusao struct {
	XMLName       xml.Name      `xml:"evtExclusao"`
	IdAttr        string        `xml:"Id,attr"`
	IdeEvento     string        `xml:"ideEvento"`
	IdeEmpregador string        `xml:"ideEmpregador"`
	InfoExclusao  *InfoExclusao `xml:"infoExclusao"`
}
