package schemes

import (
	"encoding/xml"
)

// InfoPgto ...
type InfoPgto struct {
	XMLName      xml.Name     `xml:"infoPgto"`
	DtPgto       string       `xml:"dtPgto"`
	TpPgto       int8         `xml:"tpPgto"`
	PerRef       string       `xml:"perRef"`
	IdeDmDev     string       `xml:"ideDmDev"`
	VrLiq        string       `xml:"vrLiq"`
	PaisResidExt string       `xml:"paisResidExt"`
	InfoPgtoExt  *InfoPgtoExt `xml:"infoPgtoExt"`
}
