package schemas

import (
	"encoding/xml"
)

// InfoBenInicio ...
type InfoBenInicio struct {
	XMLName        xml.Name        `xml:"infoBenInicio"`
	CadIni         string          `xml:"cadIni"`
	IndSitBenef    int8            `xml:"indSitBenef"`
	NrBeneficio    string          `xml:"nrBeneficio"`
	DtIniBeneficio string          `xml:"dtIniBeneficio"`
	DtPublic       string          `xml:"dtPublic"`
	DadosBeneficio *DadosBeneficio `xml:"dadosBeneficio"`
	SucessaoBenef  *SucessaoBenef  `xml:"sucessaoBenef"`
	MudancaCPF     *MudancaCPF     `xml:"mudancaCPF"`
	InfoBenTermino *InfoBenTermino `xml:"infoBenTermino"`
}
