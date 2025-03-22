package schema

import (
	"encoding/xml"
)

// InfoContr ...
type InfoContr struct {
	XMLName      xml.Name        `xml:"infoContr"`
	TpContr      int8            `xml:"tpContr"`
	IndContr     string          `xml:"indContr"`
	DtAdmOrig    string          `xml:"dtAdmOrig"`
	IndReint     string          `xml:"indReint"`
	IndCateg     string          `xml:"indCateg"`
	IndNatAtiv   string          `xml:"indNatAtiv"`
	IndMotDeslig string          `xml:"indMotDeslig"`
	Matricula    string          `xml:"matricula"`
	CodCateg     string          `xml:"codCateg"`
	DtInicio     string          `xml:"dtInicio"`
	InfoCompl    *InfoCompl      `xml:"infoCompl"`
	MudCategAtiv []*MudCategAtiv `xml:"mudCategAtiv"`
	UnicContr    []*UnicContr    `xml:"unicContr"`
	IdeEstab     *IdeEstab       `xml:"ideEstab"`
}
