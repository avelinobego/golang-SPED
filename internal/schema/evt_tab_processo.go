package schema

import (
	"encoding/xml"
)

// EvtTabProcesso ...
type EvtTabProcesso struct {
	XMLName       xml.Name      `xml:"evtTabProcesso"`
	IdAttr        string        `xml:"Id,attr"`
	IdeEvento     string        `xml:"ideEvento"`
	IdeEmpregador string        `xml:"ideEmpregador"`
	InfoProcesso  *InfoProcesso `xml:"infoProcesso"`
}
