package schemes

import (
	"encoding/xml"
)

// EvtAdmPrelim ...
type EvtAdmPrelim struct {
	XMLName       xml.Name       `xml:"evtAdmPrelim"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador string         `xml:"ideEmpregador"`
	InfoRegPrelim *InfoRegPrelim `xml:"infoRegPrelim"`
}
