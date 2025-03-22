package schema

import (
	"encoding/xml"
)

// DadosTrabalhador ...
type DadosTrabalhador struct {
	XMLName         xml.Name         `xml:"dadosTrabalhador"`
	NmTrab          string           `xml:"nmTrab"`
	Sexo            string           `xml:"sexo"`
	RacaCor         string           `xml:"racaCor"`
	EstCiv          string           `xml:"estCiv"`
	GrauInstr       string           `xml:"grauInstr"`
	NmSoc           string           `xml:"nmSoc"`
	PaisNac         string           `xml:"paisNac"`
	Endereco        *Endereco        `xml:"endereco"`
	TrabImig        *TrabImig        `xml:"trabImig"`
	InfoDeficiencia *InfoDeficiencia `xml:"infoDeficiencia"`
	Dependente      []*Dependente    `xml:"dependente"`
	Contato         string           `xml:"contato"`
}
