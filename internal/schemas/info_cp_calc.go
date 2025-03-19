package schemas

import (
	"encoding/xml"
)

// InfoCpCalc ...
type InfoCpCalc struct {
	XMLName   xml.Name `xml:"infoCpCalc"`
	TpCR      int      `xml:"tpCR"`
	VrCpSeg   string   `xml:"vrCpSeg"`
	VrDescSeg string   `xml:"vrDescSeg"`
}
