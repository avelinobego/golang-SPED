package schemas

import (
	"encoding/xml"
)

// InfoAnotJud ...
type InfoAnotJud struct {
	XMLName       xml.Name        `xml:"infoAnotJud"`
	CpfTrab       string          `xml:"cpfTrab"`
	NmTrab        string          `xml:"nmTrab"`
	DtNascto      string          `xml:"dtNascto"`
	DtAdm         string          `xml:"dtAdm"`
	Matricula     string          `xml:"matricula"`
	CodCateg      string          `xml:"codCateg"`
	NatAtividade  string          `xml:"natAtividade"`
	TpContr       string          `xml:"tpContr"`
	DtTerm        string          `xml:"dtTerm"`
	TpInscTrab    string          `xml:"tpInscTrab"`
	LocalTrabalho string          `xml:"localTrabalho"`
	TpRegTrab     string          `xml:"tpRegTrab"`
	TpRegPrev     string          `xml:"tpRegPrev"`
	Cargo         []*Cargo        `xml:"cargo"`
	Remuneracao   []*Remuneracao  `xml:"remuneracao"`
	Incorporacao  []*Incorporacao `xml:"incorporacao"`
	Afastamento   *Afastamento    `xml:"afastamento"`
	Desligamento  *Desligamento   `xml:"desligamento"`
}
