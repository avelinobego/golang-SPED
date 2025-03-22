package schema

import (
	"encoding/xml"
)

// InfoProcRet ...
type InfoProcRet struct {
	XMLName     xml.Name       `xml:"infoProcRet"`
	TpProcRet   string         `xml:"tpProcRet"`
	NrProcRet   string         `xml:"nrProcRet"`
	CodSusp     string         `xml:"codSusp"`
	InfoValores []*InfoValores `xml:"infoValores"`
}
