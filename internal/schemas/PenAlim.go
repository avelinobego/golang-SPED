package schemes

import (
	"encoding/xml"
)

// PenAlim ...
type PenAlim struct {
	XMLName       xml.Name `xml:"penAlim"`
	TpRend        string   `xml:"tpRend"`
	CpfDep        string   `xml:"cpfDep"`
	VlrDedPenAlim string   `xml:"vlrDedPenAlim"`
}
