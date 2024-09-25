package schemes

import (
	"encoding/xml"
)

// InfoComplementares ...
type InfoComplementares struct {
	XMLName               xml.Name               `xml:"infoComplementares"`
	CargoFuncao           *CargoFuncao           `xml:"cargoFuncao"`
	Remuneracao           string                 `xml:"remuneracao"`
	FGTS                  *FGTS                  `xml:"FGTS"`
	InfoDirigenteSindical *InfoDirigenteSindical `xml:"infoDirigenteSindical"`
	InfoTrabCedido        *InfoTrabCedido        `xml:"infoTrabCedido"`
	InfoMandElet          *InfoMandElet          `xml:"infoMandElet"`
	InfoEstagiario        string                 `xml:"infoEstagiario"`
	LocalTrabGeral        string                 `xml:"localTrabGeral"`
}
