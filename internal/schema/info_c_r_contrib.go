package schema

import (
	"encoding/xml"
)

// InfoCRContrib ...
type InfoCRContrib struct {
	XMLName xml.Name `xml:"infoCRContrib"`
	TpCR    int      `xml:"tpCR"`
	VrCR    string   `xml:"vrCR"`
}
