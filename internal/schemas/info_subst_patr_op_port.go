package schemas

import (
	"encoding/xml"
)

// InfoSubstPatrOpPort ...
type InfoSubstPatrOpPort struct {
	XMLName    xml.Name `xml:"infoSubstPatrOpPort"`
	CodLotacao string   `xml:"codLotacao"`
}
