package schemas

import (
	"encoding/xml"
)

// InfoContrato ...
type InfoContrato struct {
	XMLName        xml.Name       `xml:"infoContrato"`
	NmCargo        string         `xml:"nmCargo"`
	CBOCargo       string         `xml:"CBOCargo"`
	NmFuncao       string         `xml:"nmFuncao"`
	CBOFuncao      string         `xml:"CBOFuncao"`
	AcumCargo      string         `xml:"acumCargo"`
	CodCateg       string         `xml:"codCateg"`
	Remuneracao    string         `xml:"remuneracao"`
	Duracao        *Duracao       `xml:"duracao"`
	LocalTrabalho  *LocalTrabalho `xml:"localTrabalho"`
	HorContratual  string         `xml:"horContratual"`
	AlvaraJudicial string         `xml:"alvaraJudicial"`
	Observacoes    []*Observacoes `xml:"observacoes"`
	TreiCap        []string       `xml:"treiCap"`
}
