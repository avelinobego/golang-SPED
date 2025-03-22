package schema

import (
	"encoding/xml"
)

// InfoRRA ...
type InfoRRA struct {
	XMLName     xml.Name     `xml:"infoRRA"`
	TpProcRRA   string       `xml:"tpProcRRA"`
	NrProcRRA   string       `xml:"nrProcRRA"`
	DescRRA     string       `xml:"descRRA"`
	QtdMesesRRA string       `xml:"qtdMesesRRA"`
	DespProcJud *DespProcJud `xml:"despProcJud"`
	IdeAdv      []*IdeAdv    `xml:"ideAdv"`
}
