package schemas

import (
	"encoding/xml"
)

// DadosBeneficio ...
type DadosBeneficio struct {
	XMLName      xml.Name      `xml:"dadosBeneficio"`
	TpBeneficio  string        `xml:"tpBeneficio"`
	TpPlanRP     string        `xml:"tpPlanRP"`
	Dsc          string        `xml:"dsc"`
	IndDecJud    string        `xml:"indDecJud"`
	InfoPenMorte *InfoPenMorte `xml:"infoPenMorte"`
}
