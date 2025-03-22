package schema

import (
	"encoding/xml"
)

// TdetVerbas ...
type TdetVerbas struct {
	XMLName    xml.Name `xml:"T_detVerbas"`
	CodRubr    string   `xml:"codRubr"`
	IdeTabRubr string   `xml:"ideTabRubr"`
	QtdRubr    string   `xml:"qtdRubr"`
	FatorRubr  string   `xml:"fatorRubr"`
	VrRubr     string   `xml:"vrRubr"`
	IndApurIR  string   `xml:"indApurIR"`
}
