package schemas

import (
	"encoding/xml"
)

// InfoReintegr ...
type InfoReintegr struct {
	XMLName       xml.Name `xml:"infoReintegr"`
	TpReint       int8     `xml:"tpReint"`
	NrProcJud     string   `xml:"nrProcJud"`
	NrLeiAnistia  string   `xml:"nrLeiAnistia"`
	DtEfetRetorno string   `xml:"dtEfetRetorno"`
	DtEfeito      string   `xml:"dtEfeito"`
}
