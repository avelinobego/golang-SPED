package schemas

import (
	"encoding/xml"
)

// EvtInfoEmpregador ...
type EvtInfoEmpregador struct {
	XMLName        xml.Name        `xml:"evtInfoEmpregador"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      string          `xml:"ideEvento"`
	IdeEmpregador  *IdeEmpregador  `xml:"ideEmpregador"`
	InfoEmpregador *InfoEmpregador `xml:"infoEmpregador"`
}
