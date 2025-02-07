package schemes

import (
	"encoding/xml"
)

// FpasLotacao is Deve haver pelo menos um processo em {procJudTerceiro}(1020_infoLotacao_inclusao_dadosLotacao_fpasLotacao_infoProcJudTerceiros_procJudTerceiro) para cada código de Terceiro cujo recolhimento esteja suspenso.
type FpasLotacao struct {
	XMLName              xml.Name              `xml:"fpasLotacao"`
	Fpas                 string                `xml:"fpas"`
	CodTercs             string                `xml:"codTercs"`
	CodTercsSusp         string                `xml:"codTercsSusp"`
	InfoProcJudTerceiros *InfoProcJudTerceiros `xml:"infoProcJudTerceiros"`
}
