package schemes

import (
	"encoding/xml"
)

// InfoTercSusp ...
type InfoTercSusp struct {
	XMLName xml.Name `xml:"infoTercSusp"`
	CodTerc string   `xml:"codTerc"`
}
