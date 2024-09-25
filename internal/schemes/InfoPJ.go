package schemes

import (
	"encoding/xml"
)

// InfoPJ ...
type InfoPJ struct {
	XMLName              xml.Name    `xml:"infoPJ"`
	IndCoop              string      `xml:"indCoop"`
	IndConstr            string      `xml:"indConstr"`
	IndSubstPatr         string      `xml:"indSubstPatr"`
	PercRedContrib       string      `xml:"percRedContrib"`
	PercTransf           string      `xml:"percTransf"`
	IndTribFolhaPisPasep string      `xml:"indTribFolhaPisPasep"`
	InfoAtConc           *InfoAtConc `xml:"infoAtConc"`
}
