package schemas

import (
	"encoding/xml"
)

// InfoComplObra ...
type InfoComplObra struct {
	XMLName          xml.Name `xml:"infoComplObra"`
	IndSubstPatrObra string   `xml:"indSubstPatrObra"`
}
