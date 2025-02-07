package schemes

import (
	"encoding/xml"
)

// TidePeriodo is CHAVE_GRUPO: {iniValid*}, {fimValid*}
type TidePeriodo struct {
	XMLName  xml.Name `xml:"T_idePeriodo"`
	IniValid string   `xml:"iniValid"`
	FimValid string   `xml:"fimValid"`
}
