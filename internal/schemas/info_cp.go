package schemas

import (
	"encoding/xml"
)

// InfoCp ...
type InfoCp struct {
	XMLName     xml.Name       `xml:"infoCp"`
	ClassTrib   string         `xml:"classTrib"`
	IdeEstabLot []*IdeEstabLot `xml:"ideEstabLot"`
}
