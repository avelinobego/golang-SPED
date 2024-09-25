package schemes

import (
	"encoding/xml"
)

// TransfDom ...
type TransfDom struct {
	XMLName        xml.Name `xml:"transfDom"`
	CpfSubstituido string   `xml:"cpfSubstituido"`
	MatricAnt      string   `xml:"matricAnt"`
	DtTransf       string   `xml:"dtTransf"`
}
