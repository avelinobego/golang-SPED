package schemas

import (
	"encoding/xml"
)

// DmDev ...
type DmDev struct {
	XMLName     xml.Name       `xml:"dmDev"`
	IdeDmDev    string         `xml:"ideDmDev"`
	IndRRA      string         `xml:"indRRA"`
	InfoRRA     string         `xml:"infoRRA"`
	IdeEstabLot []*IdeEstabLot `xml:"ideEstabLot"`
}
