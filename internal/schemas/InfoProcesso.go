package schemes

import (
	"encoding/xml"
)

// InfoProcesso ...
type InfoProcesso struct {
	XMLName   xml.Name   `xml:"infoProcesso"`
	Inclusao  *Inclusao  `xml:"inclusao"`
	Alteracao *Alteracao `xml:"alteracao"`
	Exclusao  *Exclusao  `xml:"exclusao"`
}
