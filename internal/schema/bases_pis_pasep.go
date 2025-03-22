package schema

import (
	"encoding/xml"
)

// BasesPisPasep ...
type BasesPisPasep struct {
	XMLName          xml.Name `xml:"basesPisPasep"`
	VrBcPisPasep     string   `xml:"vrBcPisPasep"`
	VrBcPisPasepSusp string   `xml:"vrBcPisPasepSusp"`
}
