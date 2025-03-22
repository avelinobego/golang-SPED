package schema

import (
	"encoding/xml"
)

// DadosBenef ...
type DadosBenef struct {
	XMLName    xml.Name      `xml:"dadosBenef"`
	NmBenefic  string        `xml:"nmBenefic"`
	Sexo       string        `xml:"sexo"`
	RacaCor    string        `xml:"racaCor"`
	EstCiv     string        `xml:"estCiv"`
	IncFisMen  string        `xml:"incFisMen"`
	Endereco   *Endereco     `xml:"endereco"`
	Dependente []*Dependente `xml:"dependente"`
}
