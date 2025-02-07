package schemes

import (
	"encoding/xml"
)

// InfoBasePerAntE ...
type InfoBasePerAntE struct {
	XMLName     xml.Name       `xml:"infoBasePerAntE"`
	PerRef      string         `xml:"perRef"`
	TpAcConv    string         `xml:"tpAcConv"`
	BasePerAntE []*BasePerAntE `xml:"basePerAntE"`
}
