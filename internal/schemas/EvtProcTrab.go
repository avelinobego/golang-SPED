package schemas

import (
	"encoding/xml"
)

// EvtProcTrab ...
type EvtProcTrab struct {
	XMLName       xml.Name       `xml:"evtProcTrab"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador *IdeEmpregador `xml:"ideEmpregador"`
	InfoProcesso  *InfoProcesso  `xml:"infoProcesso"`
	IdeTrab       *IdeTrab       `xml:"ideTrab"`
}
