package schemes

import (
	"encoding/xml"
)

// TransfTit ...
type TransfTit struct {
	XMLName       xml.Name `xml:"transfTit"`
	CpfSubstituto string   `xml:"cpfSubstituto"`
	DtNascto      string   `xml:"dtNascto"`
}
