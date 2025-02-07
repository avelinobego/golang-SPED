package schemes

import (
	"encoding/xml"
)

// Medico ...
type Medico struct {
	XMLName xml.Name `xml:"medico"`
	NmMed   string   `xml:"nmMed"`
	NrCRM   string   `xml:"nrCRM"`
	UfCRM   string   `xml:"ufCRM"`
}
