package schemas

import (
	"encoding/xml"
)

// InfoInterm ...
type InfoInterm struct {
	XMLName xml.Name `xml:"infoInterm"`
	Dia     string   `xml:"dia"`
	HrsTrab string   `xml:"hrsTrab"`
}
