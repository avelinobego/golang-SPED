package schema

import (
	"encoding/xml"
)

// DadosIsencao ...
type DadosIsencao struct {
	XMLName      xml.Name `xml:"dadosIsencao"`
	IdeMinLei    string   `xml:"ideMinLei"`
	NrCertif     string   `xml:"nrCertif"`
	DtEmisCertif string   `xml:"dtEmisCertif"`
	DtVencCertif string   `xml:"dtVencCertif"`
	NrProtRenov  string   `xml:"nrProtRenov"`
	DtProtRenov  string   `xml:"dtProtRenov"`
	DtDou        string   `xml:"dtDou"`
	PagDou       int      `xml:"pagDou"`
}
