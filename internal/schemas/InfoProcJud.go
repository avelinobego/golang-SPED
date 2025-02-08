package schemas

import (
	"encoding/xml"
)

// InfoProcJud ...
type InfoProcJud struct {
	XMLName  xml.Name `xml:"infoProcJud"`
	DtSent   string   `xml:"dtSent"`
	UfVara   string   `xml:"ufVara"`
	CodMunic string   `xml:"codMunic"`
	IdVara   int      `xml:"idVara"`
}
