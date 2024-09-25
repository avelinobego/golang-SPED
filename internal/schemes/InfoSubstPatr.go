package schemes

import (
	"encoding/xml"
)

// InfoSubstPatr ...
type InfoSubstPatr struct {
	XMLName        xml.Name `xml:"infoSubstPatr"`
	IndSubstPatr   string   `xml:"indSubstPatr"`
	PercRedContrib string   `xml:"percRedContrib"`
}
