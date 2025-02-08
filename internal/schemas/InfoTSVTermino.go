package schemas

import (
	"encoding/xml"
)

// InfoTSVTermino ...
type InfoTSVTermino struct {
	XMLName       xml.Name       `xml:"infoTSVTermino"`
	DtTerm        string         `xml:"dtTerm"`
	MtvDesligTSV  string         `xml:"mtvDesligTSV"`
	PensAlim      string         `xml:"pensAlim"`
	PercAliment   string         `xml:"percAliment"`
	VrAlim        string         `xml:"vrAlim"`
	NrProcTrab    string         `xml:"nrProcTrab"`
	MudancaCPF    *MudancaCPF    `xml:"mudancaCPF"`
	VerbasResc    *VerbasResc    `xml:"verbasResc"`
	RemunAposTerm *RemunAposTerm `xml:"remunAposTerm"`
}
