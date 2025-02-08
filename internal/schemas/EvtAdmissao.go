package schemas

import (
	"encoding/xml"
)

// EvtAdmissao ...
type EvtAdmissao struct {
	XMLName       xml.Name     `xml:"evtAdmissao"`
	IdAttr        string       `xml:"Id,attr"`
	IdeEvento     string       `xml:"ideEvento"`
	IdeEmpregador string       `xml:"ideEmpregador"`
	Trabalhador   *Trabalhador `xml:"trabalhador"`
	Vinculo       *Vinculo     `xml:"vinculo"`
}
