package schemes

import (
	"encoding/xml"
)

// EvtContProc ...
type EvtContProc struct {
	XMLName       xml.Name       `xml:"evtContProc"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador *IdeEmpregador `xml:"ideEmpregador"`
	IdeProc       *IdeProc       `xml:"ideProc"`
	IdeTrab       []*IdeTrab     `xml:"ideTrab"`
}
