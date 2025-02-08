package schemas

import (
	"encoding/xml"
)

// TremunPerAnt is Validação: Deve corresponder à matrícula informada pelo empregador no evento S-2200 ou S-2300 do respectivo contrato. Não preencher no caso de Trabalhador Sem Vínculo de Emprego/Estatutário - TSVE sem informação de matrícula no evento S-2300 ou, no caso de {remunPerAnt}(1202_dmDev_infoPerAnt_idePeriodo_ideEstab_remunPerAnt), se {remunOrgSuc}(1202_dmDev_infoPerAnt_remunOrgSuc) = [S].
type TremunPerAnt struct {
	XMLName    xml.Name `xml:"T_remunPerAnt"`
	Matricula  string   `xml:"matricula"`
	ItensRemun []string `xml:"itensRemun"`
}
