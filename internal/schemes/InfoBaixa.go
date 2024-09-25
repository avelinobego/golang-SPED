package schemes

import (
	"encoding/xml"
)

// InfoBaixa ...
type InfoBaixa struct {
	XMLName      xml.Name `xml:"infoBaixa"`
	MtvDeslig    string   `xml:"mtvDeslig"`
	DtDeslig     string   `xml:"dtDeslig"`
	DtProjFimAPI string   `xml:"dtProjFimAPI"`
	NrProcTrab   string   `xml:"nrProcTrab"`
	Observacao   string   `xml:"observacao"`
}
