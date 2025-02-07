package schemes

import (
	"encoding/xml"
)

// InfoComplem ...
type InfoComplem struct {
	XMLName      xml.Name      `xml:"infoComplem"`
	NmTrab       string        `xml:"nmTrab"`
	DtNascto     string        `xml:"dtNascto"`
	SucessaoVinc *SucessaoVinc `xml:"sucessaoVinc"`
}
