package schemas

import (
	"encoding/xml"
)

// InfoEstab ...
type InfoEstab struct {
	XMLName   xml.Name   `xml:"infoEstab"`
	Inclusao  *Inclusao  `xml:"inclusao"`
	Alteracao *Alteracao `xml:"alteracao"`
	Exclusao  *Exclusao  `xml:"exclusao"`
}
