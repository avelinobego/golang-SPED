package schema

import (
	"encoding/xml"
)

// IdeProcessoPisPasep ...
type IdeProcessoPisPasep struct {
	XMLName xml.Name `xml:"ideProcessoPisPasep"`
	NrProc  string   `xml:"nrProc"`
	CodSusp string   `xml:"codSusp"`
}
