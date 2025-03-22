package schema

import (
	"encoding/xml"
)

// ProcJudTerceiro is Validação: Deve ser um código de Terceiro válido e compatível com o FPAS/Terceiros informado no grupo superior, conforme Tabela 04.
type ProcJudTerceiro struct {
	XMLName   xml.Name `xml:"procJudTerceiro"`
	CodTerc   string   `xml:"codTerc"`
	NrProcJud string   `xml:"nrProcJud"`
	CodSusp   string   `xml:"codSusp"`
}
