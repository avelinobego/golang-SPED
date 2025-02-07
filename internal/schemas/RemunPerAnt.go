package schemes

import (
	"encoding/xml"
)

// RemunPerAnt ...
type RemunPerAnt struct {
	XMLName      xml.Name `xml:"remunPerAnt"`
	Matricula    string   `xml:"matricula"`
	IndSimples   string   `xml:"indSimples"`
	ItensRemun   []string `xml:"itensRemun"`
	InfoAgNocivo string   `xml:"infoAgNocivo"`
}
