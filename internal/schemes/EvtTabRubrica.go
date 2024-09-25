package schemes

import (
	"encoding/xml"
)

// EvtTabRubrica ...
type EvtTabRubrica struct {
	XMLName       xml.Name     `xml:"evtTabRubrica"`
	IdAttr        string       `xml:"Id,attr"`
	IdeEvento     string       `xml:"ideEvento"`
	IdeEmpregador string       `xml:"ideEmpregador"`
	InfoRubrica   *InfoRubrica `xml:"infoRubrica"`
}
