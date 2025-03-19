package schemas

import (
	"encoding/xml"
)

// InfoEntEduc ...
type InfoEntEduc struct {
	XMLName xml.Name `xml:"infoEntEduc"`
	NrInsc  string   `xml:"nrInsc"`
}
