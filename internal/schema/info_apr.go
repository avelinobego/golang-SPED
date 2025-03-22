package schema

import (
	"encoding/xml"
)

// InfoApr ...
type InfoApr struct {
	XMLName     xml.Name       `xml:"infoApr"`
	NrProcJud   string         `xml:"nrProcJud"`
	InfoEntEduc []*InfoEntEduc `xml:"infoEntEduc"`
}
