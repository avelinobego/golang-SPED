package schema

import (
	"encoding/xml"
)

// Cessao ...
type Cessao struct {
	XMLName     xml.Name `xml:"cessao"`
	DtIniCessao string   `xml:"dtIniCessao"`
}
