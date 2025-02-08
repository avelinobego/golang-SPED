package schemas

import (
	"encoding/xml"
)

// InfoRubrica ...
type InfoRubrica struct {
	XMLName   xml.Name   `xml:"infoRubrica"`
	Inclusao  *Inclusao  `xml:"inclusao"`
	Alteracao *Alteracao `xml:"alteracao"`
	Exclusao  *Exclusao  `xml:"exclusao"`
}
