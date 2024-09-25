package schemes

import (
	"encoding/xml"
)

// DtAcConv is Validação: Preenchimento obrigatório e exclusivo se {classTrib}(1000_infoEmpregador_inclusao_infoCadastro_classTrib) em S-1000 = [22], {natAtividade}(./natAtividade) = [2] e {indApuracao}(1200_ideEvento_indApuracao) = [1]. Neste caso, preencher com um número entre 0 e 31, de acordo com o calendário anual.
type DtAcConv string

// TitensRemun ...
type TitensRemun struct {
	XMLName    xml.Name `xml:"T_itensRemun"`
	CodRubr    string   `xml:"codRubr"`
	IdeTabRubr string   `xml:"ideTabRubr"`
	QtdRubr    string   `xml:"qtdRubr"`
	FatorRubr  string   `xml:"fatorRubr"`
	VrRubr     string   `xml:"vrRubr"`
	IndApurIR  string   `xml:"indApurIR"`
}
