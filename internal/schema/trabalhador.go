package schema

import (
	"encoding/xml"
)

// Trabalhador ...
type Trabalhador struct {
	XMLName         xml.Name         `xml:"trabalhador"`
	CpfTrab         string           `xml:"cpfTrab"`
	NmTrab          string           `xml:"nmTrab"`
	Sexo            string           `xml:"sexo"`
	RacaCor         string           `xml:"racaCor"`
	EstCiv          string           `xml:"estCiv"`
	GrauInstr       string           `xml:"grauInstr"`
	NmSoc           string           `xml:"nmSoc"`
	Nascimento      string           `xml:"nascimento"`
	Endereco        *Endereco        `xml:"endereco"`
	TrabImig        *TrabImig        `xml:"trabImig"`
	InfoDeficiencia *InfoDeficiencia `xml:"infoDeficiencia"`
	Dependente      []*Dependente    `xml:"dependente"`
	Contato         string           `xml:"contato"`
}
