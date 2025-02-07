package schemes

import (
	"encoding/xml"
)

// InfoBenAlteracao ...
type InfoBenAlteracao struct {
	XMLName        xml.Name        `xml:"infoBenAlteracao"`
	DtAltBeneficio string          `xml:"dtAltBeneficio"`
	DadosBeneficio *DadosBeneficio `xml:"dadosBeneficio"`
}
