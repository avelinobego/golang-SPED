package schema

import (
	"encoding/xml"
)

// TinfoCadastro is CONDICAO_GRUPO: N (se {procEmi}(1000_ideEvento_procEmi) = [8]); O (nos demais casos).
type TinfoCadastro struct {
	XMLName              xml.Name              `xml:"T_infoCadastro"`
	ClassTrib            string                `xml:"classTrib"`
	IndCoop              string                `xml:"indCoop"`
	IndConstr            string                `xml:"indConstr"`
	IndDesFolha          int8                  `xml:"indDesFolha"`
	IndOpcCP             int8                  `xml:"indOpcCP"`
	IndPorte             string                `xml:"indPorte"`
	IndOptRegEletron     int8                  `xml:"indOptRegEletron"`
	CnpjEFR              string                `xml:"cnpjEFR"`
	DtTrans11096         string                `xml:"dtTrans11096"`
	IndTribFolhaPisPasep string                `xml:"indTribFolhaPisPasep"`
	DadosIsencao         *DadosIsencao         `xml:"dadosIsencao"`
	InfoOrgInternacional *InfoOrgInternacional `xml:"infoOrgInternacional"`
}
