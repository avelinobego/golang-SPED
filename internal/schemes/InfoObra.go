package schemes

import (
	"encoding/xml"
)

// InfoObra ...
type InfoObra struct {
	XMLName          xml.Name `xml:"infoObra"`
	IndSubstPatrObra string   `xml:"indSubstPatrObra"`
}
