package schemas

import (
	"encoding/xml"
)

// EvtAfastTemp ...
type EvtAfastTemp struct {
	XMLName         xml.Name         `xml:"evtAfastTemp"`
	IdAttr          string           `xml:"Id,attr"`
	IdeEvento       string           `xml:"ideEvento"`
	IdeEmpregador   string           `xml:"ideEmpregador"`
	IdeVinculo      *IdeVinculo      `xml:"ideVinculo"`
	InfoAfastamento *InfoAfastamento `xml:"infoAfastamento"`
}
