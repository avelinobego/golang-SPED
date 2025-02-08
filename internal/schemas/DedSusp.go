package schemas

import (
	"encoding/xml"
)

// DedSusp ...
type DedSusp struct {
	XMLName       xml.Name    `xml:"dedSusp"`
	IndTpDeducao  string      `xml:"indTpDeducao"`
	VlrDedSusp    string      `xml:"vlrDedSusp"`
	CnpjEntidPC   string      `xml:"cnpjEntidPC"`
	VlrPatrocFunp string      `xml:"vlrPatrocFunp"`
	BenefPen      []*BenefPen `xml:"benefPen"`
}
