package schemes

import (
	"encoding/xml"
)

// InfoComProd ...
type InfoComProd struct {
	XMLName    xml.Name    `xml:"infoComProd"`
	IdeEstabel *IdeEstabel `xml:"ideEstabel"`
}
