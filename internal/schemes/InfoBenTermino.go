package schemes

import (
	"encoding/xml"
)

// InfoBenTermino ...
type InfoBenTermino struct {
	XMLName         xml.Name `xml:"infoBenTermino"`
	DtTermBeneficio string   `xml:"dtTermBeneficio"`
	MtvTermino      string   `xml:"mtvTermino"`
	CnpjOrgaoSuc    string   `xml:"cnpjOrgaoSuc"`
	NovoCPF         string   `xml:"novoCPF"`
}
