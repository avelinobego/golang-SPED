package schemes

import (
	"encoding/xml"
)

// RemunPerApur ...
type RemunPerApur struct {
	XMLName      xml.Name `xml:"remunPerApur"`
	Matricula    string   `xml:"matricula"`
	IndSimples   string   `xml:"indSimples"`
	ItensRemun   []string `xml:"itensRemun"`
	InfoAgNocivo string   `xml:"infoAgNocivo"`
}
