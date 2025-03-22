package schema

import (
	"encoding/xml"
)

// InfoTerm ...
type InfoTerm struct {
	XMLName      xml.Name `xml:"infoTerm"`
	DtTerm       string   `xml:"dtTerm"`
	MtvDesligTSV string   `xml:"mtvDesligTSV"`
}
