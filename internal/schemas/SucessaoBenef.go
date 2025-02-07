package schemes

import (
	"encoding/xml"
)

// SucessaoBenef ...
type SucessaoBenef struct {
	XMLName        xml.Name `xml:"sucessaoBenef"`
	CnpjOrgaoAnt   string   `xml:"cnpjOrgaoAnt"`
	NrBeneficioAnt string   `xml:"nrBeneficioAnt"`
	DtTransf       string   `xml:"dtTransf"`
	Observacao     string   `xml:"observacao"`
}
