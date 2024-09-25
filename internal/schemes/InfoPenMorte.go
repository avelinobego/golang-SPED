package schemes

import (
	"encoding/xml"
)

// InfoPenMorte ...
type InfoPenMorte struct {
	XMLName      xml.Name      `xml:"infoPenMorte"`
	TpPenMorte   string        `xml:"tpPenMorte"`
	InstPenMorte *InstPenMorte `xml:"instPenMorte"`
}
