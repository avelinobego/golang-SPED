package schemas

import (
	"encoding/xml"
)

// DadosOpPort ...
type DadosOpPort struct {
	XMLName xml.Name `xml:"dadosOpPort"`
	AliqRat string   `xml:"aliqRat"`
	Fap     string   `xml:"fap"`
}
