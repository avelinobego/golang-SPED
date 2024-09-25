package schemes

import (
	"encoding/xml"
)

// IdeLocalAcid ...
type IdeLocalAcid struct {
	XMLName xml.Name `xml:"ideLocalAcid"`
	TpInsc  string   `xml:"tpInsc"`
	NrInsc  string   `xml:"nrInsc"`
}
