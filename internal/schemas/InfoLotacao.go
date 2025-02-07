package schemes

import (
	"encoding/xml"
)

// InfoLotacao ...
type InfoLotacao struct {
	XMLName   xml.Name   `xml:"infoLotacao"`
	Inclusao  *Inclusao  `xml:"inclusao"`
	Alteracao *Alteracao `xml:"alteracao"`
	Exclusao  *Exclusao  `xml:"exclusao"`
}
