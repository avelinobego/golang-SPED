package schemes

import (
	"encoding/xml"
)

// InfoBaseCS ...
type InfoBaseCS struct {
	XMLName xml.Name `xml:"infoBaseCS"`
	Ind13   string   `xml:"ind13"`
	TpValor int8     `xml:"tpValor"`
	Valor   string   `xml:"valor"`
}
