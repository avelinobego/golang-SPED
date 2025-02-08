package schemas

import (
	"encoding/xml"
)

// TideRubrica is REGRA:REGRA_CARACTERE_ESPECIAL
type TideRubrica struct {
	XMLName    xml.Name `xml:"T_ideRubrica"`
	CodRubr    string   `xml:"codRubr"`
	IdeTabRubr string   `xml:"ideTabRubr"`
	IniValid   string   `xml:"iniValid"`
	FimValid   string   `xml:"fimValid"`
}
