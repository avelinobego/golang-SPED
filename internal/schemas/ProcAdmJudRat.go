package schemas

import (
	"encoding/xml"
)

// ProcAdmJudRat is Validação: Deve ser um número de processo administrativo ou judicial válido e existente na Tabela de Processos (S-1070), com {indMatProc}(1070_infoProcesso_inclusao_dadosProc_indMatProc) = [1].
type ProcAdmJudRat struct {
	XMLName xml.Name `xml:"procAdmJudRat"`
	TpProc  string   `xml:"tpProc"`
	NrProc  string   `xml:"nrProc"`
	CodSusp string   `xml:"codSusp"`
}
