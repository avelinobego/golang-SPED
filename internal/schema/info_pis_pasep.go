package schema

import (
	"encoding/xml"
)

// InfoPisPasep ...
type InfoPisPasep struct {
	XMLName  xml.Name    `xml:"infoPisPasep"`
	IdeEstab []*IdeEstab `xml:"ideEstab"`
}
