package schema

import (
	"encoding/xml"
)

// InfoDep ...
type InfoDep struct {
	XMLName  xml.Name `xml:"infoDep"`
	CpfDep   string   `xml:"cpfDep"`
	DtNascto string   `xml:"dtNascto"`
	Nome     string   `xml:"nome"`
	DepIRRF  string   `xml:"depIRRF"`
	TpDep    string   `xml:"tpDep"`
	DescrDep string   `xml:"descrDep"`
}
