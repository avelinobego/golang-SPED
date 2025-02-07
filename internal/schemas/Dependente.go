package schemes

import (
	"encoding/xml"
)

// Dependente ...
type Dependente struct {
	XMLName  xml.Name `xml:"dependente"`
	TpDep    string   `xml:"tpDep"`
	NmDep    string   `xml:"nmDep"`
	DtNascto string   `xml:"dtNascto"`
	CpfDep   string   `xml:"cpfDep"`
	DepIRRF  string   `xml:"depIRRF"`
	DepSF    string   `xml:"depSF"`
	IncTrab  string   `xml:"incTrab"`
	DescrDep string   `xml:"descrDep"`
}
