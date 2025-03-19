package schemas

import (
	"encoding/xml"
)

// IdeProcessoCP ...
type IdeProcessoCP struct {
	XMLName    xml.Name `xml:"ideProcessoCP"`
	TpProc     string   `xml:"tpProc"`
	NrProc     string   `xml:"nrProc"`
	ExtDecisao int8     `xml:"extDecisao"`
	CodSusp    string   `xml:"codSusp"`
}
