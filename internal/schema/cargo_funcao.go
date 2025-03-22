package schema

import (
	"encoding/xml"
)

// CargoFuncao ...
type CargoFuncao struct {
	XMLName   xml.Name `xml:"cargoFuncao"`
	NmCargo   string   `xml:"nmCargo"`
	CBOCargo  string   `xml:"CBOCargo"`
	NmFuncao  string   `xml:"nmFuncao"`
	CBOFuncao string   `xml:"CBOFuncao"`
}
