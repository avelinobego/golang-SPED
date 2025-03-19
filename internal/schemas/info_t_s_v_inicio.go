package schemas

import (
	"encoding/xml"
)

// InfoTSVInicio ...
type InfoTSVInicio struct {
	XMLName            xml.Name            `xml:"infoTSVInicio"`
	CadIni             string              `xml:"cadIni"`
	Matricula          string              `xml:"matricula"`
	CodCateg           string              `xml:"codCateg"`
	DtInicio           string              `xml:"dtInicio"`
	NrProcTrab         string              `xml:"nrProcTrab"`
	NatAtividade       string              `xml:"natAtividade"`
	InfoComplementares *InfoComplementares `xml:"infoComplementares"`
	MudancaCPF         *MudancaCPF         `xml:"mudancaCPF"`
	Afastamento        *Afastamento        `xml:"afastamento"`
	Termino            *Termino            `xml:"termino"`
}
