package schemes

import (
	"encoding/xml"
)

// IdePeriodo ...
type IdePeriodo struct {
	XMLName  xml.Name    `xml:"idePeriodo"`
	PerRef   string      `xml:"perRef"`
	IdeEstab []*IdeEstab `xml:"ideEstab"`
}