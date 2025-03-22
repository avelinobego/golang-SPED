package schema

import (
	"encoding/xml"
)

// IdeADC ...
type IdeADC struct {
	XMLName    xml.Name      `xml:"ideADC"`
	DtAcConv   string        `xml:"dtAcConv"`
	TpAcConv   string        `xml:"tpAcConv"`
	Dsc        string        `xml:"dsc"`
	RemunSuc   string        `xml:"remunSuc"`
	IdePeriodo []*IdePeriodo `xml:"idePeriodo"`
}
