package schemes

import (
	"encoding/xml"
)

// InfoPerApur ...
type InfoPerApur struct {
	XMLName  xml.Name    `xml:"infoPerApur"`
	IdeEstab []*IdeEstab `xml:"ideEstab"`
}
