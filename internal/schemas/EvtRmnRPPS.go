package schemas

import (
	"encoding/xml"
)

// EvtRmnRPPS ...
type EvtRmnRPPS struct {
	XMLName        xml.Name        `xml:"evtRmnRPPS"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      string          `xml:"ideEvento"`
	IdeEmpregador  string          `xml:"ideEmpregador"`
	IdeTrabalhador *IdeTrabalhador `xml:"ideTrabalhador"`
	DmDev          []*DmDev        `xml:"dmDev"`
}
