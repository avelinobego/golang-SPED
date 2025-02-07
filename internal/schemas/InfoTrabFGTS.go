package schemes

import (
	"encoding/xml"
)

// InfoTrabFGTS ...
type InfoTrabFGTS struct {
	XMLName          xml.Name          `xml:"infoTrabFGTS"`
	Matricula        string            `xml:"matricula"`
	CodCateg         string            `xml:"codCateg"`
	CategOrig        string            `xml:"categOrig"`
	InfoFGTSProcTrab *InfoFGTSProcTrab `xml:"infoFGTSProcTrab"`
}
