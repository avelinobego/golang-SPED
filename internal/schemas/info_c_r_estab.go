package schemas

import (
	"encoding/xml"
)

// InfoCREstab ...
type InfoCREstab struct {
	XMLName  xml.Name `xml:"infoCREstab"`
	TpCR     int      `xml:"tpCR"`
	VrCR     string   `xml:"vrCR"`
	VrSuspCR string   `xml:"vrSuspCR"`
}
