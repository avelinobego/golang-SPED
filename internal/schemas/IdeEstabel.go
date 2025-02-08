package schemas

import (
	"encoding/xml"
)

// IdeEstabel ...
type IdeEstabel struct {
	XMLName          xml.Name    `xml:"ideEstabel"`
	NrInscEstabRural string      `xml:"nrInscEstabRural"`
	TpComerc         []*TpComerc `xml:"tpComerc"`
}
