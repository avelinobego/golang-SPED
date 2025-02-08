package schemas

import (
	"encoding/xml"
)

// EvtTabLotacao ...
type EvtTabLotacao struct {
	XMLName       xml.Name     `xml:"evtTabLotacao"`
	IdAttr        string       `xml:"Id,attr"`
	IdeEvento     string       `xml:"ideEvento"`
	IdeEmpregador string       `xml:"ideEmpregador"`
	InfoLotacao   *InfoLotacao `xml:"infoLotacao"`
}
