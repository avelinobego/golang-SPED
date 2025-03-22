package schema

import (
	"encoding/xml"
)

// TideEstabperAnt is DESCRICAO_COMPLETA:Rubricas que compõem o provento ou pensão do beneficiário.
type TideEstabperAnt struct {
	XMLName    xml.Name `xml:"T_ideEstab_perAnt"`
	ItensRemun []string `xml:"itensRemun"`
	*TideEstab
}
