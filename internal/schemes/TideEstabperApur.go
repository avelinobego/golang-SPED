package schemes

import (
	"encoding/xml"
)

// TideEstabperApur is DESCRICAO_COMPLETA:Rubricas que compõem o provento ou pensão do beneficiário.
type TideEstabperApur struct {
	XMLName    xml.Name `xml:"T_ideEstab_perApur"`
	ItensRemun []string `xml:"itensRemun"`
	*TideEstab
}
