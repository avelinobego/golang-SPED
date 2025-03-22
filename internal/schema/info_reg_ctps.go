package schema

import (
	"encoding/xml"
)

// InfoRegCTPS ...
type InfoRegCTPS struct {
	XMLName    xml.Name `xml:"infoRegCTPS"`
	CBOCargo   string   `xml:"CBOCargo"`
	VrSalFx    string   `xml:"vrSalFx"`
	UndSalFixo string   `xml:"undSalFixo"`
	TpContr    string   `xml:"tpContr"`
	DtTerm     string   `xml:"dtTerm"`
}
