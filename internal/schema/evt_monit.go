package schema

import (
	"encoding/xml"
)

// EvtMonit ...
type EvtMonit struct {
	XMLName       xml.Name   `xml:"evtMonit"`
	IdAttr        string     `xml:"Id,attr"`
	IdeEvento     string     `xml:"ideEvento"`
	IdeEmpregador string     `xml:"ideEmpregador"`
	IdeVinculo    string     `xml:"ideVinculo"`
	ExMedOcup     *ExMedOcup `xml:"exMedOcup"`
}
