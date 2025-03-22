package schema

import (
	"encoding/xml"
)

// EvtPgtos ...
type EvtPgtos struct {
	XMLName       xml.Name  `xml:"evtPgtos"`
	IdAttr        string    `xml:"Id,attr"`
	IdeEvento     string    `xml:"ideEvento"`
	IdeEmpregador string    `xml:"ideEmpregador"`
	IdeBenef      *IdeBenef `xml:"ideBenef"`
}
