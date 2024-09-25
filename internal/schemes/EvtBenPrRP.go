package schemes

import (
	"encoding/xml"
)

// EvtBenPrRP ...
type EvtBenPrRP struct {
	XMLName       xml.Name  `xml:"evtBenPrRP"`
	IdAttr        string    `xml:"Id,attr"`
	IdeEvento     string    `xml:"ideEvento"`
	IdeEmpregador string    `xml:"ideEmpregador"`
	IdeBenef      *IdeBenef `xml:"ideBenef"`
	DmDev         []*DmDev  `xml:"dmDev"`
}
