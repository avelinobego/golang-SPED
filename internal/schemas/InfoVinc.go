package schemas

import (
	"encoding/xml"
)

// InfoVinc ...
type InfoVinc struct {
	XMLName      xml.Name       `xml:"infoVinc"`
	TpRegTrab    string         `xml:"tpRegTrab"`
	TpRegPrev    string         `xml:"tpRegPrev"`
	DtAdm        string         `xml:"dtAdm"`
	TmpParc      string         `xml:"tmpParc"`
	Duracao      *Duracao       `xml:"duracao"`
	Observacoes  []*Observacoes `xml:"observacoes"`
	SucessaoVinc *SucessaoVinc  `xml:"sucessaoVinc"`
	InfoDeslig   *InfoDeslig    `xml:"infoDeslig"`
}
