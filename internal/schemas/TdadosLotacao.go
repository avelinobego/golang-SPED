package schemas

import (
	"encoding/xml"
)

// TdadosLotacao is b) Deve ser um identificador válido, constante das bases da RFB.
type TdadosLotacao struct {
	XMLName         xml.Name         `xml:"T_dadosLotacao"`
	TpLotacao       string           `xml:"tpLotacao"`
	TpInsc          string           `xml:"tpInsc"`
	NrInsc          string           `xml:"nrInsc"`
	FpasLotacao     *FpasLotacao     `xml:"fpasLotacao"`
	InfoEmprParcial *InfoEmprParcial `xml:"infoEmprParcial"`
	DadosOpPort     *DadosOpPort     `xml:"dadosOpPort"`
}
