package schemas

import (
	"encoding/xml"
)

// InfoVlr ...
type InfoVlr struct {
	XMLName    xml.Name      `xml:"infoVlr"`
	CompIni    string        `xml:"compIni"`
	CompFim    string        `xml:"compFim"`
	IndReperc  int8          `xml:"indReperc"`
	IndenSD    string        `xml:"indenSD"`
	IndenAbono string        `xml:"indenAbono"`
	Abono      []*Abono      `xml:"abono"`
	IdePeriodo []*IdePeriodo `xml:"idePeriodo"`
}
