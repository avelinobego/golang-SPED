package schema

import (
	"encoding/xml"
)

// InfoCategPisPasep ...
type InfoCategPisPasep struct {
	XMLName          xml.Name            `xml:"infoCategPisPasep"`
	Matricula        string              `xml:"matricula"`
	CodCateg         string              `xml:"codCateg"`
	InfoBasePisPasep []*InfoBasePisPasep `xml:"infoBasePisPasep"`
}
