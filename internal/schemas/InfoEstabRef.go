package schemas

import (
	"encoding/xml"
)

// InfoEstabRef ...
type InfoEstabRef struct {
	XMLName      xml.Name `xml:"infoEstabRef"`
	AliqRat      string   `xml:"aliqRat"`
	Fap          string   `xml:"fap"`
	AliqRatAjust float64  `xml:"aliqRatAjust"`
}
