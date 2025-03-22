package schema

import (
	"encoding/xml"
)

// IdeDep ...
type IdeDep struct {
	XMLName  xml.Name `xml:"ideDep"`
	CpfDep   string   `xml:"cpfDep"`
	DepIRRF  string   `xml:"depIRRF"`
	DtNascto string   `xml:"dtNascto"`
	Nome     string   `xml:"nome"`
	TpDep    string   `xml:"tpDep"`
	DescrDep string   `xml:"descrDep"`
}
