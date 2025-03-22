package schema

import (
	"encoding/xml"
)

// EvtBasesTrab ...
type EvtBasesTrab struct {
	XMLName        xml.Name        `xml:"evtBasesTrab"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      string          `xml:"ideEvento"`
	IdeEmpregador  string          `xml:"ideEmpregador"`
	IdeTrabalhador *IdeTrabalhador `xml:"ideTrabalhador"`
	InfoCpCalc     []*InfoCpCalc   `xml:"infoCpCalc"`
	InfoCp         *InfoCp         `xml:"infoCp"`
	InfoPisPasep   *InfoPisPasep   `xml:"infoPisPasep"`
}
