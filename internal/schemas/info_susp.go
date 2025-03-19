package schemas

import (
	"encoding/xml"
)

// InfoSusp ...
type InfoSusp struct {
	XMLName     xml.Name `xml:"infoSusp"`
	CodSusp     string   `xml:"codSusp"`
	IndSusp     string   `xml:"indSusp"`
	DtDecisao   string   `xml:"dtDecisao"`
	IndDeposito string   `xml:"indDeposito"`
}
