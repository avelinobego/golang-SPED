package schemas

import (
	"encoding/xml"
)

// InfoBaseFGTS ...
type InfoBaseFGTS struct {
	XMLName         xml.Name           `xml:"infoBaseFGTS"`
	BasePerApur     []*BasePerApur     `xml:"basePerApur"`
	InfoBasePerAntE []*InfoBasePerAntE `xml:"infoBasePerAntE"`
}
