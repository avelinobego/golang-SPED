package schema

import (
	"encoding/xml"
)

// EvtRemun ...
type EvtRemun struct {
	XMLName        xml.Name        `xml:"evtRemun"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      string          `xml:"ideEvento"`
	IdeEmpregador  string          `xml:"ideEmpregador"`
	IdeTrabalhador *IdeTrabalhador `xml:"ideTrabalhador"`
	DmDev          []*DmDev        `xml:"dmDev"`
}
