package schemas

import (
	"encoding/xml"
)

// TideEstab is Validação: Deve ser compatível com o conteúdo do campo {ideEstab/tpInsc}(./tpInsc). Deve ser um identificador válido, constante das bases da RFB, vinculado ao empregador.
type TideEstab struct {
	XMLName  xml.Name `xml:"T_ideEstab"`
	TpInsc   string   `xml:"tpInsc"`
	NrInsc   string   `xml:"nrInsc"`
	IniValid string   `xml:"iniValid"`
	FimValid string   `xml:"fimValid"`
}
