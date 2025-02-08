package schemas

import (
	"encoding/xml"
)

// IdeProc ...
type IdeProc struct {
	XMLName      xml.Name        `xml:"ideProc"`
	NrProcTrab   string          `xml:"nrProcTrab"`
	PerApur      string          `xml:"perApur"`
	InfoTributos []*InfoTributos `xml:"infoTributos"`
	InfoCRIRRF   []*InfoCRIRRF   `xml:"infoCRIRRF"`
}
