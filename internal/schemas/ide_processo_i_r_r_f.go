package schemas

import (
	"encoding/xml"
)

// IdeProcessoIRRF ...
type IdeProcessoIRRF struct {
	XMLName xml.Name `xml:"ideProcessoIRRF"`
	NrProc  string   `xml:"nrProc"`
	CodSusp string   `xml:"codSusp"`
}
