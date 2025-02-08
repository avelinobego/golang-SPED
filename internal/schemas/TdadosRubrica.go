package schemas

import (
	"encoding/xml"
)

// TdadosRubrica is Informar a descrição (nome) da rubrica no sistema de folha de pagamento da empresa.
type TdadosRubrica struct {
	XMLName             xml.Name               `xml:"T_dadosRubrica"`
	DscRubr             string                 `xml:"dscRubr"`
	NatRubr             int                    `xml:"natRubr"`
	TpRubr              int8                   `xml:"tpRubr"`
	CodIncCP            string                 `xml:"codIncCP"`
	CodIncIRRF          int                    `xml:"codIncIRRF"`
	CodIncFGTS          string                 `xml:"codIncFGTS"`
	CodIncCPRP          string                 `xml:"codIncCPRP"`
	CodIncPisPasep      string                 `xml:"codIncPisPasep"`
	TetoRemun           string                 `xml:"tetoRemun"`
	Observacao          string                 `xml:"observacao"`
	IdeProcessoCP       []*IdeProcessoCP       `xml:"ideProcessoCP"`
	IdeProcessoIRRF     []*IdeProcessoIRRF     `xml:"ideProcessoIRRF"`
	IdeProcessoFGTS     []*IdeProcessoFGTS     `xml:"ideProcessoFGTS"`
	IdeProcessoPisPasep []*IdeProcessoPisPasep `xml:"ideProcessoPisPasep"`
}
