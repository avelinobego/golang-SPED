package schema

import (
	"encoding/xml"
)

// TideLotacao is REGRA:REGRA_CARACTERE_ESPECIAL
type TideLotacao struct {
	XMLName    xml.Name `xml:"T_ideLotacao"`
	CodLotacao string   `xml:"codLotacao"`
	IniValid   string   `xml:"iniValid"`
	FimValid   string   `xml:"fimValid"`
}
