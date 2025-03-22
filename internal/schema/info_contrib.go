package schema

import (
	"encoding/xml"
)

// InfoContrib ...
type InfoContrib struct {
	XMLName   xml.Name `xml:"infoContrib"`
	ClassTrib string   `xml:"classTrib"`
	InfoPJ    *InfoPJ  `xml:"infoPJ"`
}
