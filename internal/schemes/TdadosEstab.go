package schemes

import (
	"encoding/xml"
)

// TdadosEstab is Deve ser um identificador válido, constante das bases da RFB, vinculado ao empregador.
type TdadosEstab struct {
	XMLName    xml.Name    `xml:"T_dadosEstab"`
	CnaePrep   string      `xml:"cnaePrep"`
	CnpjResp   string      `xml:"cnpjResp"`
	AliqGilrat *AliqGilrat `xml:"aliqGilrat"`
	InfoCaepf  *InfoCaepf  `xml:"infoCaepf"`
	InfoObra   *InfoObra   `xml:"infoObra"`
	InfoTrab   *InfoTrab   `xml:"infoTrab"`
}
