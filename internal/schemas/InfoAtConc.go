package schemas

import (
	"encoding/xml"
)

// InfoAtConc ...
type InfoAtConc struct {
	XMLName  xml.Name `xml:"infoAtConc"`
	FatorMes string   `xml:"fatorMes"`
	Fator13  string   `xml:"fator13"`
}
