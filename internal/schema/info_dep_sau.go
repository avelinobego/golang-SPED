package schema

import (
	"encoding/xml"
)

// InfoDepSau ...
type InfoDepSau struct {
	XMLName     xml.Name `xml:"infoDepSau"`
	CpfDep      string   `xml:"cpfDep"`
	VlrSaudeDep string   `xml:"vlrSaudeDep"`
}
