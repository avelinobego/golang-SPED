package schemes

import (
	"encoding/xml"
)

// PerAquis ...
type PerAquis struct {
	XMLName  xml.Name `xml:"perAquis"`
	DtInicio string   `xml:"dtInicio"`
	DtFim    string   `xml:"dtFim"`
}
