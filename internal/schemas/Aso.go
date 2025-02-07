package schemes

import (
	"encoding/xml"
)

// Aso ...
type Aso struct {
	XMLName xml.Name `xml:"aso"`
	DtAso   string   `xml:"dtAso"`
	ResAso  int8     `xml:"resAso"`
	Exame   []*Exame `xml:"exame"`
	Medico  *Medico  `xml:"medico"`
}
