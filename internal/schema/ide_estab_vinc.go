package schema

import (
	"encoding/xml"
)

// IdeEstabVinc ...
type IdeEstabVinc struct {
	XMLName xml.Name `xml:"ideEstabVinc"`
	TpInsc  string   `xml:"tpInsc"`
	NrInsc  string   `xml:"nrInsc"`
}
