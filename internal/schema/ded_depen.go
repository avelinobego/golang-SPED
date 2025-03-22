package schema

import (
	"encoding/xml"
)

// DedDepen ...
type DedDepen struct {
	XMLName   xml.Name `xml:"dedDepen"`
	TpRend    string   `xml:"tpRend"`
	CpfDep    string   `xml:"cpfDep"`
	VlrDedDep string   `xml:"vlrDedDep"`
}
