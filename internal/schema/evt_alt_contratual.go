package schema

import (
	"encoding/xml"
)

// EvtAltContratual ...
type EvtAltContratual struct {
	XMLName       xml.Name       `xml:"evtAltContratual"`
	IdAttr        string         `xml:"Id,attr"`
	IdeEvento     string         `xml:"ideEvento"`
	IdeEmpregador string         `xml:"ideEmpregador"`
	IdeVinculo    string         `xml:"ideVinculo"`
	AltContratual *AltContratual `xml:"altContratual"`
}
