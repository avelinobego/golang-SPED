package schemas

import (
	"encoding/xml"
)

// InfoEmpregador ...
type InfoEmpregador struct {
	XMLName   xml.Name   `xml:"infoEmpregador"`
	Inclusao  *Inclusao  `xml:"inclusao"`
	Alteracao *Alteracao `xml:"alteracao"`
	Exclusao  *Exclusao  `xml:"exclusao"`
}
