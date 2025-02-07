package schemes

import (
	"encoding/xml"
)

// IdeAdv ...
type IdeAdv struct {
	XMLName xml.Name `xml:"ideAdv"`
	TpInsc  string   `xml:"tpInsc"`
	NrInsc  string   `xml:"nrInsc"`
	VlrAdv  string   `xml:"vlrAdv"`
}
