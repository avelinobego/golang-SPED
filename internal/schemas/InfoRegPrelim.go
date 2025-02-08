package schemas

import (
	"encoding/xml"
)

// InfoRegPrelim ...
type InfoRegPrelim struct {
	XMLName      xml.Name     `xml:"infoRegPrelim"`
	CpfTrab      string       `xml:"cpfTrab"`
	DtNascto     string       `xml:"dtNascto"`
	DtAdm        string       `xml:"dtAdm"`
	Matricula    string       `xml:"matricula"`
	CodCateg     string       `xml:"codCateg"`
	NatAtividade string       `xml:"natAtividade"`
	InfoRegCTPS  *InfoRegCTPS `xml:"infoRegCTPS"`
}
