package schema

import (
	"encoding/xml"
)

// EvtCAT ...
type EvtCAT struct {
	XMLName       xml.Name `xml:"evtCAT"`
	IdAttr        string   `xml:"Id,attr"`
	IdeEvento     string   `xml:"ideEvento"`
	IdeEmpregador string   `xml:"ideEmpregador"`
	IdeVinculo    string   `xml:"ideVinculo"`
	Cat           *Cat     `xml:"cat"`
}
