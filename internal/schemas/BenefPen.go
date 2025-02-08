package schemas

import (
	"encoding/xml"
)

// BenefPen ...
type BenefPen struct {
	XMLName      xml.Name `xml:"benefPen"`
	CpfDep       string   `xml:"cpfDep"`
	VlrDepenSusp string   `xml:"vlrDepenSusp"`
}
