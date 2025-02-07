package schemes

import (
	"encoding/xml"
)

// EvtAltCadastral ...
type EvtAltCadastral struct {
	XMLName        xml.Name        `xml:"evtAltCadastral"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      string          `xml:"ideEvento"`
	IdeEmpregador  string          `xml:"ideEmpregador"`
	IdeTrabalhador *IdeTrabalhador `xml:"ideTrabalhador"`
	Alteracao      *Alteracao      `xml:"alteracao"`
}
