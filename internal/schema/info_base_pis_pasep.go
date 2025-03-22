package schema

import (
	"encoding/xml"
)

// InfoBasePisPasep ...
type InfoBasePisPasep struct {
	XMLName         xml.Name `xml:"infoBasePisPasep"`
	Ind13           int8     `xml:"ind13"`
	TpValorPisPasep int8     `xml:"tpValorPisPasep"`
	ValorPisPasep   string   `xml:"valorPisPasep"`
}
