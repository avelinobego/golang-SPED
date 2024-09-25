package schemes

import (
	"encoding/xml"
)

// EvtTribProcTrab ...
type EvtTribProcTrab struct {
	XMLName       xml.Name       `xml:"evtTribProcTrab"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     *IdeEvento     `xml:"ideEvento"`
	IdeEmpregador *IdeEmpregador `xml:"ideEmpregador"`
	IdeProc       *IdeProc       `xml:"ideProc"`
}
