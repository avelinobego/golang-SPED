package schema

import (
	"encoding/xml"
)

// IdeProcessoFGTS ...
type IdeProcessoFGTS struct {
	XMLName xml.Name `xml:"ideProcessoFGTS"`
	NrProc  string   `xml:"nrProc"`
}
