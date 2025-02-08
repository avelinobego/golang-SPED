package schemas

import (
	"encoding/xml"
)

// ConsigFGTS ...
type ConsigFGTS struct {
	XMLName   xml.Name `xml:"consigFGTS"`
	InsConsig string   `xml:"insConsig"`
	NrContr   string   `xml:"nrContr"`
}
