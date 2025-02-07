package schemes

import (
	"encoding/xml"
)

// EvtExcProcTrab ...
type EvtExcProcTrab struct {
	XMLName       xml.Name      `xml:"evtExcProcTrab"`
	IdAttr        string        `xml:"Id,attr"`
	IdeEvento     string        `xml:"ideEvento"`
	IdeEmpregador string        `xml:"ideEmpregador"`
	InfoExclusao  *InfoExclusao `xml:"infoExclusao"`
}
