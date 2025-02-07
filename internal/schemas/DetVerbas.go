package schemes

import (
	"encoding/xml"
)

// DetVerbas ...
type DetVerbas struct {
	XMLName    xml.Name `xml:"detVerbas"`
	CodRubr    string   `xml:"codRubr"`
	IdeTabRubr string   `xml:"ideTabRubr"`
	QtdRubr    string   `xml:"qtdRubr"`
	FatorRubr  string   `xml:"fatorRubr"`
	VrRubr     string   `xml:"vrRubr"`
	IndApurIR  string   `xml:"indApurIR"`
	DescFolha  string   `xml:"descFolha"`
}
