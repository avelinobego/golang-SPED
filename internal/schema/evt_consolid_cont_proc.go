package schema

import (
	"encoding/xml"
)

// EvtConsolidContProc ...
type EvtConsolidContProc struct {
	XMLName       xml.Name       `xml:"evtConsolidContProc"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador *IdeEmpregador `xml:"ideEmpregador"`
	IdeProc       *IdeProc       `xml:"ideProc"`
}
