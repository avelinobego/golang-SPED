package schema

import (
	"encoding/xml"
)

// EvtComProd ...
type EvtComProd struct {
	XMLName       xml.Name       `xml:"evtComProd"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador *IdeEmpregador `xml:"ideEmpregador"`
	InfoComProd   *InfoComProd   `xml:"infoComProd"`
}
