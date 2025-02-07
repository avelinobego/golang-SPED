package schemes

import (
	"encoding/xml"
)

// EvtToxic ...
type EvtToxic struct {
	XMLName       xml.Name      `xml:"evtToxic"`
	IdAttr        string        `xml:"Id,attr"`
	IdeEvento     string        `xml:"ideEvento"`
	IdeEmpregador string        `xml:"ideEmpregador"`
	IdeVinculo    *IdeVinculo   `xml:"ideVinculo"`
	Toxicologico  *Toxicologico `xml:"toxicologico"`
}
