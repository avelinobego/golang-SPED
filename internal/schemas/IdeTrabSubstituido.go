package schemes

import (
	"encoding/xml"
)

// IdeTrabSubstituido ...
type IdeTrabSubstituido struct {
	XMLName      xml.Name `xml:"ideTrabSubstituido"`
	CpfTrabSubst string   `xml:"cpfTrabSubst"`
}
