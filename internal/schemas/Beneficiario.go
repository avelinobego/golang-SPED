package schemas

import (
	"encoding/xml"
)

// Beneficiario ...
type Beneficiario struct {
	XMLName    xml.Name `xml:"beneficiario"`
	CpfBenef   string   `xml:"cpfBenef"`
	Matricula  string   `xml:"matricula"`
	CnpjOrigem string   `xml:"cnpjOrigem"`
}
