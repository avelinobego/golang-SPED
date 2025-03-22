package schema

import (
	"encoding/xml"
)

// EConsignado ...
type EConsignado struct {
	XMLName       xml.Name `xml:"eConsignado"`
	InstFinanc    string   `xml:"instFinanc"`
	NrContrato    string   `xml:"nrContrato"`
	VreConsignado string   `xml:"vreConsignado"`
}
