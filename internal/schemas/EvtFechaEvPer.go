package schemas

import (
	"encoding/xml"
)

// EvtFechaEvPer ...
type EvtFechaEvPer struct {
	XMLName       xml.Name  `xml:"evtFechaEvPer"`
	IdAttr        string    `xml:"Id,attr"`
	IdeEvento     string    `xml:"ideEvento"`
	IdeEmpregador string    `xml:"ideEmpregador"`
	InfoFech      *InfoFech `xml:"infoFech"`
}
