package schemes

import (
	"encoding/xml"
)

// InfoFech ...
type InfoFech struct {
	XMLName         xml.Name `xml:"infoFech"`
	EvtRemun        string   `xml:"evtRemun"`
	EvtPgtos        string   `xml:"evtPgtos"`
	EvtComProd      string   `xml:"evtComProd"`
	EvtContratAvNP  string   `xml:"evtContratAvNP"`
	EvtInfoComplPer string   `xml:"evtInfoComplPer"`
	IndExcApur1250  string   `xml:"indExcApur1250"`
	TransDCTFWeb    string   `xml:"transDCTFWeb"`
	NaoValid        string   `xml:"naoValid"`
}
