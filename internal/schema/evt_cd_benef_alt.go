package schema

import (
	"encoding/xml"
)

// EvtCdBenefAlt ...
type EvtCdBenefAlt struct {
	XMLName       xml.Name   `xml:"evtCdBenefAlt"`
	IdAttr        string     `xml:"Id,attr"`
	IdeEvento     string     `xml:"ideEvento"`
	IdeEmpregador string     `xml:"ideEmpregador"`
	IdeBenef      *IdeBenef  `xml:"ideBenef"`
	Alteracao     *Alteracao `xml:"alteracao"`
}
