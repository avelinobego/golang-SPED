package schemes

import (
	"encoding/xml"
)

// IniAfastamento ...
type IniAfastamento struct {
	XMLName        xml.Name      `xml:"iniAfastamento"`
	DtIniAfast     string        `xml:"dtIniAfast"`
	CodMotAfast    string        `xml:"codMotAfast"`
	InfoMesmoMtv   string        `xml:"infoMesmoMtv"`
	TpAcidTransito int8          `xml:"tpAcidTransito"`
	Observacao     string        `xml:"observacao"`
	PerAquis       *PerAquis     `xml:"perAquis"`
	InfoCessao     *InfoCessao   `xml:"infoCessao"`
	InfoMandSind   *InfoMandSind `xml:"infoMandSind"`
	InfoMandElet   *InfoMandElet `xml:"infoMandElet"`
}
