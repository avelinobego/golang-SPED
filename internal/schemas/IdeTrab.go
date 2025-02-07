package schemes

import (
	"encoding/xml"
)

// IdeTrab ...
type IdeTrab struct {
	XMLName   xml.Name     `xml:"ideTrab"`
	CpfTrab   string       `xml:"cpfTrab"`
	NmTrab    string       `xml:"nmTrab"`
	DtNascto  string       `xml:"dtNascto"`
	InfoContr []*InfoContr `xml:"infoContr"`
}
