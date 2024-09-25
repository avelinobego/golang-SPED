package schemes

import (
	"encoding/xml"
)

// InfoIRCR ...
type InfoIRCR struct {
	XMLName     xml.Name       `xml:"infoIRCR"`
	TpCR        string         `xml:"tpCR"`
	DedDepen    []*DedDepen    `xml:"dedDepen"`
	PenAlim     []*PenAlim     `xml:"penAlim"`
	PrevidCompl []*PrevidCompl `xml:"previdCompl"`
	InfoProcRet []*InfoProcRet `xml:"infoProcRet"`
}
