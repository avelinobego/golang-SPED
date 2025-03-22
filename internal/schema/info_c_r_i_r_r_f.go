package schema

import (
	"encoding/xml"
)

// InfoCRIRRF ...
type InfoCRIRRF struct {
	XMLName xml.Name `xml:"infoCRIRRF"`
	TpCR    int      `xml:"tpCR"`
	VrCR    string   `xml:"vrCR"`
}
