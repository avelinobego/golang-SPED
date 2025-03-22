package schema

import (
	"encoding/xml"
)

// InfoProcJudRub ...
type InfoProcJudRub struct {
	XMLName  xml.Name `xml:"infoProcJudRub"`
	NrProc   string   `xml:"nrProc"`
	UfVara   string   `xml:"ufVara"`
	CodMunic string   `xml:"codMunic"`
	IdVara   int      `xml:"idVara"`
}
