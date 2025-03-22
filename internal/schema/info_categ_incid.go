package schema

import (
	"encoding/xml"
)

// InfoCategIncid ...
type InfoCategIncid struct {
	XMLName    xml.Name      `xml:"infoCategIncid"`
	Matricula  string        `xml:"matricula"`
	CodCateg   string        `xml:"codCateg"`
	IndSimples string        `xml:"indSimples"`
	InfoBaseCS []*InfoBaseCS `xml:"infoBaseCS"`
	CalcTerc   []*CalcTerc   `xml:"calcTerc"`
	InfoPerRef []*InfoPerRef `xml:"infoPerRef"`
}
