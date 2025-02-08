package schemas

import (
	"encoding/xml"
)

// ExMedOcup ...
type ExMedOcup struct {
	XMLName     xml.Name   `xml:"exMedOcup"`
	TpExameOcup int8       `xml:"tpExameOcup"`
	Aso         *Aso       `xml:"aso"`
	RespMonit   *RespMonit `xml:"respMonit"`
}
